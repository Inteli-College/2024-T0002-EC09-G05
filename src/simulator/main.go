package main

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"math/rand"
	"time"
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

func configureClient(clientName string) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("rabbitmq:1883")
	opts.SetClientID(clientName)
	opts.SetUsername("guest")
	opts.SetPassword("guest")

	opts.OnConnect = func(client mqtt.Client) {
		fmt.Println("Connected")
	}

	opts.OnConnectionLost = func(client mqtt.Client, err error) {
		fmt.Printf("Connect lost: %v", err)
	}

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return client
}

func simulateSensor(clientName string) {

	client := configureClient(clientName)

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

func main() {
	go simulateSensor("Liberdade")
	go simulateSensor("Luz")
	go simulateSensor("Butantã")
	go simulateSensor("Morumbi")
	go simulateSensor("Pinheiros")
	go simulateSensor("Ipiranga")
	go simulateSensor("Parelheiros")
	go simulateSensor("Grajaú")
	go simulateSensor("Moema")
	go simulateSensor("Vila Mariana")
	go simulateSensor("Paulista")
	select {}
}
