package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Sensor struct {
	nameType        string
	min             float32
	max             float32
	connection      bool
	counter         int
	status          string
	startWave       bool
	currentPosition float32
	goPosition      float32
	waveSize        int
}

type SensorData struct {
	SensorId  string  `json:"sensor_id"`
	Timestamp int64   `json:"timestamp"`
	PM1_0     float32 `json:"pm1_0"`
	PM2_5     float32 `json:"pm2_5"`
	PM4_0     float32 `json:"pm4_0"`
	PM10_0    float32 `json:"pm10_0"`
	Temp      float32 `json:"temp"`
	Hum       float32 `json:"hum"`
}

var PM1_0 = NewSensor("PM1_0", 0, 37)
var PM2_5 = NewSensor("PM2_5", 0, 30)
var PM4_0 = NewSensor("PM4_0", -40, 70)
var PM10_0 = NewSensor("PM10_0", 3, 30)
var Temp = NewSensor("Temp", -40, 70)
var Hum = NewSensor("Hum", 20, 50)

func configureSSL() *tls.Config {
	certpool := x509.NewCertPool()
	ca, err := os.ReadFile("/app/certs/cert.pem") // CA cert
	if err != nil {
		fmt.Printf("Erro ao ler arquivo de certificado: %s\n", err)
		os.Exit(1)
	}
	certpool.AppendCertsFromPEM(ca)

	// Carrega o certificado e a chave do cliente
	cert, err := tls.LoadX509KeyPair("/app/certs/client.crt", "/app/certs/client.key")
	if err != nil {
		fmt.Printf("Erro ao carregar o par de certificado/chave do cliente: %s\n", err)
		os.Exit(1)
	}

	return &tls.Config{
		RootCAs:      certpool,
		Certificates: []tls.Certificate{cert},
	}
}

func NewSensor(name string, min float32, max float32) *Sensor {
	return &Sensor{
		nameType:        name,
		min:             min,
		max:             max,
		connection:      false,
		counter:         10,
		status:          "new wave",
		startWave:       true,
		currentPosition: 0,
		goPosition:      0,
		waveSize:        0,
	}
}

func (s *Sensor) generateWaveData(currentPosition float32) float32 {
	var value float32

	if s.status == "new wave" {
		s.waveSize = 5
		if s.startWave {
			s.currentPosition = currentPosition
			s.startWave = false
			value = s.currentPosition

		}
		s.status = "wave"
		//fmt.Println("Start wave")

	} else if s.status == "wave" {
		//fmt.Println("Wave")
		s.waveSize -= 1
		loss := float32(math.Cos(float64(time.Now().Unix())) * 2)
		fmt.Printf("Valor do loss %.2f", loss)
		value = s.currentPosition + loss
		if s.waveSize <= 0 {
			s.status = "transtion"
			s.goPosition = float32(rand.Intn(int(s.max-s.min)) + int(s.min))
		}
	} else if s.status == "transtion" {

		var posOrNeg float32 = 0.0
		var difference float64 = float64(s.currentPosition - s.goPosition)

		if difference > 0 {
			posOrNeg = 0.1
		} else {
			posOrNeg = -0.1
		}

		s.currentPosition -= posOrNeg
		value = s.currentPosition

		if math.Abs(difference) >= 0.10 && math.Abs(difference) <= 1.25 {
			s.currentPosition = s.goPosition
			s.status = "new wave"
		}
	} else {
		fmt.Println("Erro")
	}

	if value > s.max {
		value = s.max
	}
	if value < s.min {
		value = s.min
	}

	return value
}

func generateSensorData(lastData SensorData) SensorData {

	return SensorData{
		SensorId:  lastData.SensorId,
		Timestamp: time.Now().Unix(),
		PM1_0:     PM1_0.generateWaveData(lastData.PM1_0),
		PM2_5:     PM2_5.generateWaveData(lastData.PM2_5),
		PM4_0:     PM4_0.generateWaveData(lastData.PM4_0),
		PM10_0:    PM10_0.generateWaveData(lastData.PM10_0),
		Temp:      Temp.generateWaveData(lastData.Temp),
		Hum:       Hum.generateWaveData(lastData.Hum),
	}
}

func configureClient(certificate *tls.Config, clientName string) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("ssl://rabbitmq:8883")
	opts.SetClientID(clientName)
	opts.SetUsername("publisherUser")
	opts.SetPassword("publisherPassword")
	opts.SetTLSConfig(certificate)

	opts.OnConnect = func(client mqtt.Client) {
		fmt.Println("Connected as " + clientName)
	}

	opts.OnConnectionLost = func(client mqtt.Client, err error) {
		fmt.Printf(clientName, "-> Connect lost: %v", err)
	}

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return client
}

func simulateSensor(certificate *tls.Config, clientName string) {

	client := configureClient(certificate, clientName)

	currentData := SensorData{
		SensorId:  clientName,
		Timestamp: time.Now().Unix(),
		PM1_0:     float32(rand.Intn(int(PM1_0.max-PM1_0.min)) + int(PM1_0.min)),
		PM2_5:     float32(rand.Intn(int(PM2_5.max-PM2_5.min)) + int(PM2_5.min)),
		PM4_0:     float32(rand.Intn(int(PM4_0.max-PM4_0.min)) + int(PM4_0.min)),
		PM10_0:    float32(rand.Intn(int(PM10_0.max-PM10_0.min)) + int(PM10_0.min)),
		Temp:      float32(rand.Intn(int(Temp.max-Temp.min)) + int(Temp.min)),
		Hum:       float32(rand.Intn(int(Hum.max-Hum.min)) + int(Hum.min)),
	}

	lastData := currentData

	for {

		time.Sleep(time.Duration(1 * time.Second))

		currentData = generateSensorData(lastData)

		lastData = currentData

		jsonData, _ := json.Marshal(currentData)
		//fmt.Println(string(jsonData))

		if token := client.Publish("sensor/data", 1, true, jsonData); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

	}
}

func main() {
	sslCertificate := configureSSL()

	go simulateSensor(sslCertificate, "Liberdade")
	go simulateSensor(sslCertificate, "Luz")
	go simulateSensor(sslCertificate, "Butantã")
	go simulateSensor(sslCertificate, "Morumbi")
	go simulateSensor(sslCertificate, "Pinheiros")
	go simulateSensor(sslCertificate, "Ipiranga")
	go simulateSensor(sslCertificate, "Parelheiros")
	go simulateSensor(sslCertificate, "Grajaú")
	go simulateSensor(sslCertificate, "Moema")
	go simulateSensor(sslCertificate, "Vila Mariana")
	go simulateSensor(sslCertificate, "Paulista")
	select {}
}
