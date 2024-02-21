package tests

import (
	"context"
	"fmt"
	influx "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"strings"
	"testing"
	"time"
)

func setupInflux() (api.WriteAPIBlocking, influx.Client) {

	org := "iot"
	bucket := "sensor"

	client := influx.NewClientWithOptions("http://influxdb:8086", "admin",
		influx.DefaultOptions().SetBatchSize(1000).SetUseGZip(true).SetPrecision(time.Second))

	organization, err := client.OrganizationsAPI().FindOrganizationByName(context.Background(), "iot")
	if err != nil {
		fmt.Println("Erro ao buscar a organizacao")
		fmt.Println(err.Error())
	}

	_, err = client.BucketsAPI().FindBucketByName(context.Background(), "sensor")

	if err != nil && strings.Contains(err.Error(), "not found") {

		fmt.Printf("Erro ao buscar o bucket '%s', tentando criar um novo\n", bucket)
		_, err := client.BucketsAPI().CreateBucketWithName(context.Background(), organization, bucket)

		if err != nil {
			fmt.Println("Erro ao criar o bucket")
			fmt.Println(err.Error())

		} else {
			fmt.Printf("Bucket %s criado com sucesso\n", bucket)
		}

	} else {
		fmt.Printf("Bucket '%s' encontrado, continuando...\n", bucket)
	}

	return client.WriteAPIBlocking(org, bucket), client
}

func TestInfluxDB(t *testing.T) {
	// Setup específico do teste
	_, influxClient := setupInflux()

	t.Run("Testar conexão", func(t *testing.T) {
		_, err := influxClient.Health(context.Background())
		if err != nil {
			t.Fatalf("Erro ao conectar ao InfluxDB: %v", err)
		}
	})
}
