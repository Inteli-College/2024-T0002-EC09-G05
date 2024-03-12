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

func CreateinfluxConnection() api.QueryAPI {

	client := influx.NewClientWithOptions(influxDBURL, influxDBToken,
		influx.DefaultOptions().SetBatchSize(1000).SetUseGZip(true).SetPrecision(time.Second))

	organization, err := client.OrganizationsAPI().FindOrganizationByName(context.Background(), influxDBOrg)
	if err != nil {
		fmt.Println("Erro ao buscar a organizacao")
		fmt.Println(err.Error())
	}

	_, err = client.BucketsAPI().FindBucketByName(context.Background(), influxDBBucket)

	if err != nil && strings.Contains(err.Error(), "not found") {

		fmt.Printf("Erro ao buscar o bucket '%s', tentando criar um novo\n", influxDBBucket)
		_, err := client.BucketsAPI().CreateBucketWithName(context.Background(), organization, influxDBBucket)

		if err != nil {
			fmt.Println("Erro ao criar o bucket")
			fmt.Println(err.Error())
			os.Exit(1)

		} else {
			fmt.Printf("Bucket %s criado com sucesso\n", influxDBBucket)
		}

	} else {
		fmt.Printf("Bucket '%s' encontrado, continuando...\n", influxDBBucket)
	}

	return client.QueryAPI(influxDBOrg)
}
