package db

import (
	"context"
	"fmt"

	"github.com/influxdata/influxdb-client-go/v2/api"
)

type Measurement struct {
	Time     string  `json:"time"`
	Value    float32 `json:"value"`
}

type RelativeResponseJsonObj struct {
	Field  string        `json:"field"`
	Sensor string        `json:"sensor"`
	Data   []Measurement `json:"data"`
}

// PEGA OS DADOS COM BASE NO TEMPO RELATIVO, POR EXEMPLO, "ULTIMOS 30 MINUTOS" OU "ENTRE 2 HORAS ATRAS E AGORA"
func GetDataByRelativeTime(
	QueryAPI api.QueryAPI,
	startTime int,
	endTime *int,
	aggregator int,
	Fields []string,
	Sensors []string,
) ([]RelativeResponseJsonObj, error) {



	// Montando a string de range:
	rangeString := ""
	if endTime == nil {
		rangeString = fmt.Sprintf("|> range(start: -%dm)", startTime)
	} else {
		rangeString = fmt.Sprintf("|> range(start: -%dm, stop: -%dm)", startTime, *endTime)
	}
	// Montando a string de agregação:
	aggregatorString :=
		fmt.Sprintf("|> aggregateWindow(every: %ds, fn: mean, createEmpty: false)", aggregator)

	response := []RelativeResponseJsonObj{}

	for _, sensor := range Sensors {
		for _, field := range Fields {

			// Montando a string de campos e sensores.:
			fieldString := fmt.Sprintf(`|> filter(fn: (r) =>  r["_field"] == "%s")`, field)
			sensorString := fmt.Sprintf(`|> filter(fn: (r) =>  r["sensor_id"] == "%s")`, sensor)

			// Montando a query:
			queryString := 
				`from(bucket:"sensor")` +
				rangeString + 
				`|> filter(fn: (r) => r["_measurement"] == "medicoes")` +
				fieldString + sensorString + aggregatorString


			data, err := QueryAPI.Query(context.Background(), queryString)

			if err != nil {
				return nil, err
			}

			records := []Measurement{}

			for data.Next() {

				measure := Measurement{
					Time:     data.Record().Time().String(),
					Value:    float32(int(data.Record().Value().(float64)*100)) / 100,
				}

				records = append(records, measure)
			}

			response = append(response, RelativeResponseJsonObj{
				Field:  field,
				Sensor: sensor,
				Data:   records,
			})

		}
	}

	if len(response) > 2000 {
		fmt.Println("[WARNING] Resposta ultrapasou 2000 entradas, truncando para as ultimas 2000")
		fmt.Println("tamanho original:" + fmt.Sprint(len(response)))
		return response[:2000], nil
	}

	return response, nil

}

// Pega os dados com base em TIMESTAMPS, por exemplo, "ENTRE 22/06/2021 12:00" E "22/06/2021 13:00"
func GetDataByTimestamp(QueryAPI api.QueryAPI) {
}

func GetAllFields(QueryAPI api.QueryAPI) ([]string, error) {
	data, err := QueryAPI.Query(context.Background(), `

	from(bucket: "sensor")
	|> range(start: -1y)
	|> filter(fn: (r) => r["_measurement"] == "medicoes")
	|> keep(columns: ["_field"])
	|> distinct(column: "_field")
  
	`)

	if err != nil {
		return nil, err
	}

	result := []string{}

	for data.Next() {
		result = append(result, data.Record().ValueByKey("_field").(string))
	}
	return result, nil

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