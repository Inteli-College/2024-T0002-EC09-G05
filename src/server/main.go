package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	influx "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	amqp "github.com/rabbitmq/amqp091-go"
)

type SensorData struct {
	Sensor_id string  `json:"sensor_id"`
	Timestamp int64   `json:"timestamp"`
	PM1_0     float32 `json:"pm1_0"`
	PM2_5     float32 `json:"pm2_5"`
	PM4_0     float32 `json:"pm4_0"`
	PM10_0    float32 `json:"pm10_0"`
	Temp      float32 `json:"temp"`
	Hum       float32 `json:"hum"`
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
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		fmt.Println("ERRO NA OPERACAO DE WRITE")
		fmt.Println(err.Error())
	}
}

func setupInflux() api.WriteAPIBlocking {

	org := "iot"
	bucket := "sensor"

	client := influx.NewClientWithOptions("http://influxdb:8086", "admin",
		influx.DefaultOptions().SetBatchSize(1000).SetUseGZip(true).SetPrecision(time.Second))

	organization, _ := client.OrganizationsAPI().FindOrganizationByName(context.Background(), "iot")

	_, err := client.BucketsAPI().FindBucketByName(context.Background(), "sensor")
	if err != nil && err.Error() != "bucket 'sensor' not found"{

		fmt.Println("Erro ao buscar o bucket, tentando criar um novo")
		_, err := client.BucketsAPI().CreateBucketWithName(context.Background(), organization, bucket)
		if err != nil {
			fmt.Println("Erro ao criar o bucket")
			fmt.Println(err.Error())
		}

	} else {
		fmt.Println("Bucket encontrado, continuando...")
	}


	return client.WriteAPIBlocking(org, bucket)
}

func connectRabbitMQ() <-chan amqp.Delivery {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		defer conn.Close()
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		defer ch.Close()
	}

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

	messages, _ := ch.Consume(
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

	return messages
}

func handleSensorData( rabbitMessages <-chan amqp.Delivery, influxWriteAPI api.WriteAPIBlocking) {
	batchCounter := 0

	contagemIDs := make(map[string]int)

	fmt.Println("Waiting for messages...")

	for d := range rabbitMessages {
		
		os.Stdout.Sync()
		var data SensorData
		err := json.Unmarshal(d.Body, &data)

		contagemIDs[data.Sensor_id]++

		if err != nil {
			fmt.Println("Error parsing JSON: ", err)
		}
		saveDataToInfluxDB(data, influxWriteAPI)
		
		batchCounter++
		if batchCounter == 1000 {
			fmt.Println("--> Flushed batch (1000 points) to InfluxDB")
			fmt.Println("	", contagemIDs)
			batchCounter = 0
		} else if batchCounter%100 == 0 {
			fmt.Printf("Filling batch... (%d/1000)\n", batchCounter)
		}
	}
}

func main() {

	influxWriteAPI := setupInflux()

	rabbitMessages := connectRabbitMQ()

	go handleSensorData(rabbitMessages, influxWriteAPI)

	select {}
	
}
