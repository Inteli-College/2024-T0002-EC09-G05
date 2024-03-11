package db

import (
	"os"
	"fmt"
	"time"
	"strings"
	"context"
	"github.com/influxdata/influxdb-client-go/v2/api"
    influx "github.com/influxdata/influxdb-client-go/v2"
)

const (
    influxDBToken  = "admin"
    influxDBURL    = "http://influxdb:8086"
    influxDBOrg    = "iot"
    influxDBBucket = "sensor"
)

func SetupInflux() api.WriteAPIBlocking {

	org := "iot"
	bucket := "sensor"

	client := influx.NewClientWithOptions(influxDBURL, influxDBToken,
		influx.DefaultOptions().SetBatchSize(1000).SetUseGZip(true).SetPrecision(time.Second))

	organization, err := client.OrganizationsAPI().FindOrganizationByName(context.Background(), influxDBToken)
	if err != nil {
		fmt.Println("Erro ao buscar a organizacao")
		fmt.Println(err.Error())
	}

	_, err = client.BucketsAPI().FindBucketByName(context.Background(), influxDBBucket)

	if err != nil && strings.Contains(err.Error(), "not found") {

		fmt.Printf("Erro ao buscar o bucket '%s', tentando criar um novo\n", bucket)
		_, err := client.BucketsAPI().CreateBucketWithName(context.Background(), organization, bucket)

		if err != nil {
			fmt.Println("Erro ao criar o bucket")
			fmt.Println(err.Error())
			os.Exit(1)

		} else {
			fmt.Printf("Bucket %s criado com sucesso\n", bucket)
		}

	} else {
		fmt.Printf("Bucket '%s' encontrado, continuando...\n", bucket)
	}

	return client.WriteAPIBlocking(org, bucket)
}
