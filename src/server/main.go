package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	influx "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"context"
	"time"
	"encoding/json"
)

type SensorData struct {
	Sensor_id string  `json:"sensor_id"`
	Timestamp int64   `json:"timestamp"`
	PM1_0     float32 `json:"pm1_0"`
	PM2_5     float32 `json:"pm2_5"`
	PM4_0     float32 `json:"pm4_0"`
	PM10_0    float32 `json:"pm10_0"`
	Temp      int     `json:"temp"`
	Hum       int     `json:"hum"`
}


func saveDataToInfluxDB(data SensorData, writeAPI api.WriteAPIBlocking) {
	p := influx.NewPointWithMeasurement("medicoes").
			AddTag("sensor_id", data.Sensor_id).
			AddField("hum", data.Hum).
			AddField("temp", data.Temp).
			AddField("pm2_5", data.PM2_5).
			AddField("pm1_0", data.PM1_0).
			AddField("pm4_0", data.PM4_0).
			AddField("pm10_0", data.PM10_0).
			SetTime(time.Unix(data.Timestamp, 0))
	
		// Escreve o ponto
		writeAPI.WritePoint(context.Background(), p)
}

func main() {

	client := influx.NewClient("http://influxdb:8086", "admin")
	org := "iot"
	bucket := "sensor"
	writeAPI := client.WriteAPIBlocking(org, bucket)

	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"MQTTQueue", // name of the queue
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}
	
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-acknowledge
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf("Received a message: %s\n", d.Body)
			var data SensorData
			err := json.Unmarshal(d.Body, &data)
			if err != nil {
				fmt.Println("Error parsing JSON: ", err)
			}
			saveDataToInfluxDB(data, writeAPI)
		}
	}()

	fmt.Println(" [*] Waiting for messages")
	<-forever
}
