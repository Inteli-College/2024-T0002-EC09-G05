package db

import (
	"context"
	"fmt"

	"github.com/influxdata/influxdb-client-go/v2/api"
)

type Measurement struct {
	Time     string  `json:"time"`
	SensorID string  `json:"sensor_id"`
	Field    string  `json:"field"`
	Value    float32 `json:"value"`
}

// PEGA OS DADOS COM BASE NO TEMPO RELATIVO, POR EXEMPLO, "ULTIMOS 30 MINUTOS" OU "ENTRE 2 HORAS ATRAS E AGORA"
func GetDataByRelativeTime(
	QueryAPI api.QueryAPI,
	startTime int,
	endTime *int,
	aggregator int,
	Field []string,
) ([]Measurement, error) {

	// Montando a string de campos:
	fieldString := `|> filter(fn: (r) => `
	for _, f := range Field {
		fieldString += fmt.Sprintf(` r["_field"] == "%s" or `, f)
	}
	fieldString = fieldString[:len(fieldString)-4] + ")"

	// Montando a string de range:
	rangeString := ""
	if endTime == nil {
	rangeString =
		fmt.Sprintf("|> range(start: -%dm)", startTime)
	} else {
		rangeString = fmt.Sprintf("|> range(start: -%dm, stop: -%dm)", startTime, *endTime)
	}
	// Montando a string de agregação:
	aggregatorString :=
		fmt.Sprintf("|> aggregateWindow(every: %ds, fn: mean, createEmpty: false)", aggregator)

	// Montando a query:
	queryString := `
	from(bucket:"sensor")
	` + rangeString + `
	|> filter(fn: (r) => r["_measurement"] == "medicoes")
	` + fieldString + `
	|> filter(fn: (r) => r["sensor_id"] == "Morumbi")
	` + aggregatorString

	fmt.Println(queryString)

	data, err := QueryAPI.Query(context.Background(), queryString)

	if err != nil {
		return nil, err
	}

	response := []Measurement{}

	for data.Next() {

		measure := Measurement{
			Time:     data.Record().Time().String(),
			SensorID: data.Record().ValueByKey("sensor_id").(string),
			Field:    data.Record().Field(),
			Value:    float32(int(data.Record().Value().(float64)*100)) / 100,
		}

		response = append(response, measure)
	}

	if len(response) > 1000 {
		fmt.Println("[WARNING] Resposta ultrapasou 1000 entradas, truncando para as ultimas 1000")
		fmt.Println("tamanho original:"+fmt.Sprint(len(response)))
		return response[:1000], nil
	} 

	return response, nil

}

// Pega os dados com base em TIMESTAMPS, por exemplo, "ENTRE 22/06/2021 12:00" E "22/06/2021 13:00"
func GetDataByTimestamp(QueryAPI api.QueryAPI) {
}

func GetAllSensors(QueryAPI api.QueryAPI) ([]string, error) {
	data, err := QueryAPI.Query(context.Background(), `

	from(bucket: "sensor")
	|> range(start: -1y)
	|> filter(fn: (r) => r["_measurement"] == "medicoes")
	|> keep(columns: ["sensor_id"])
	|> distinct(column: "sensor_id")
  
	`)

	if err != nil {
		return nil, err
	}

	result := []string{}

	for data.Next() {
		result = append(result, data.Record().ValueByKey("sensor_id").(string))
	}
	return result, nil

}
