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

func generateSensorData(lastData SensorData) SensorData {
	return SensorData{
		SensorId:  lastData.SensorId,
		Timestamp: time.Now().Unix(),
		PM1_0:     lastData.PM1_0 + rand.Float32()*1.5 - 0.75,
		PM2_5:     lastData.PM2_5 + rand.Float32() - 0.5,
		PM4_0:     lastData.PM4_0 + rand.Float32() - 0.5,
		PM10_0:    lastData.PM10_0 + rand.Float32() - 0.5,
		Temp:      lastData.Temp + rand.Float32() - 0.5,
		Hum:       lastData.Hum + rand.Float32() - 0.5,
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
		PM1_0:     100,
		PM2_5:     75,
		PM4_0:     75,
		PM10_0:    50,
		Temp:      20,
		Hum:       40,
	}

	lastData := currentData

	for {

		time.Sleep(time.Duration(1 * time.Second))

		currentData = generateSensorData(lastData)

		lastData = currentData

		jsonData, _ := json.Marshal(currentData)

		if token := client.Publish("sensor/data", 1, true, jsonData); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

	}
}

type Sensor struct {
	nameType        string
	delta           float64
	min             float64
	max             float64
	connection      bool
	counter         int
	status          string
	startWave       bool
	currentPosition float64
	goPosition      float64
	waveSize        int
}

func NewSensor(name string, min float64, max float64) *Sensor {
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

func (s *Sensor) GenerateWaveData() string {
	var value float64

	if s.status == "new wave" {
		s.waveSize = 5
		if s.startWave {
			s.currentPosition = float64(rand.Intn(int(s.max-s.min)) + int(s.min))
			s.startWave = false
			value = s.currentPosition
		}
		s.status = "wave"
	} else if s.status == "wave" {
		s.waveSize -= 1
		value = s.currentPosition + math.Round(math.Abs(rand.NormFloat64()*5)*math.Cos(s.delta)*10)/10
		if s.waveSize <= 0 {
			s.status = "transtion"
			s.goPosition = float64(rand.Intn(int(s.max-s.min)) + int(s.min))
		}
	} else if s.status == "transtion" {
		fmt.Printf("Indo de %.2f para %.2f\n", s.currentPosition, s.goPosition)
		posOrNeg := 0.0
		difference := s.currentPosition - s.goPosition

		if difference > 0 {
			posOrNeg = 1.0
		} else {
			posOrNeg = -1.0
		}

		fmt.Println(difference, posOrNeg)
		s.currentPosition -= posOrNeg
		value = s.currentPosition
		if s.currentPosition == s.goPosition {
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

	return fmt.Sprintf("%.2f", value)
}

func main() {
	// sslCertificate := configureSSL()
	s := NewSensor("sensor", -10, 10)
	fmt.Fprintf(os.Stdout, "Valores do Sensor Simulado: %s\n", s.GenerateWaveData())
	// go simulateSensor(sslCertificate, "Liberdade")
	// go simulateSensor(sslCertificate, "Luz")
	// go simulateSensor(sslCertificate, "Butantã")
	// go simulateSensor(sslCertificate, "Morumbi")
	// go simulateSensor(sslCertificate, "Pinheiros")
	// go simulateSensor(sslCertificate, "Ipiranga")
	// go simulateSensor(sslCertificate, "Parelheiros")
	// go simulateSensor(sslCertificate, "Grajaú")
	// go simulateSensor(sslCertificate, "Moema")
	// go simulateSensor(sslCertificate, "Vila Mariana")
	// go simulateSensor(sslCertificate, "Paulista")
	// select {}
}
