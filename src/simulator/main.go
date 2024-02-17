package main

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"math/rand"
	"time"
)

type SensorData struct {
	Sensor_id string  `json:"sensor_id"`
	Timestamp int     `json:"timestamp"`
	PM1_0     float32 `json:"pm1_0"`
	PM2_5     float32 `json:"pm2_5"`
	PM4_0     float32 `json:"pm4_0"`
	PM10_0    float32 `json:"pm10_0"`
	Temp      int     `json:"temp"`
	Hum       int     `json:"hum"`
}

func simulateSensorData() SensorData {
	return SensorData{
		Sensor_id: "sensor1",
		Timestamp: int(time.Now().Unix()),
		PM1_0:     rand.Float32() * 50,
		PM2_5:     rand.Float32() * 50,
		PM4_0:     rand.Float32() * 50,
		PM10_0:    rand.Float32() * 50,
		Temp:      rand.Intn(120) - 20,
		Hum:       rand.Intn(120) - 20,
	}
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func main() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("rabbitmq:1883")
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername("guest")
	opts.SetPassword("guest")
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for i := 0; i < 200; i++ {

		topic := "sensor/data"
		data, _ := json.Marshal(simulateSensorData())

		if token := client.Publish(topic, 1, true, data); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

		fmt.Println("Sent simulated sensor reading to topic", topic)
		time.Sleep(3 * time.Second)
	}
	client.Disconnect(250)
}
