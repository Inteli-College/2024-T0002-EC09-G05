package tests

import (
	"fmt"
	"math/rand"
	"testing"
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

func TestGenerateSensorData(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	lastData := SensorData{
		SensorId:  "TestSensor",
		Timestamp: time.Now().Unix() - 100,
		PM1_0:     10,
		PM2_5:     20,
		PM4_0:     30,
		PM10_0:    40,
		Temp:      25,
		Hum:       50,
	}

	newData := generateSensorData(lastData)

	if newData.SensorId != lastData.SensorId {
		t.Errorf("Id do sensor se manteve. Esperado: %s, Recebido: %s", lastData.SensorId, newData.SensorId)
	}

	if newData.Timestamp <= lastData.Timestamp {
		t.Errorf("O timestamp deve ser maior que o anterior. Último registro: %d, Novo registro: %d", lastData.Timestamp, newData.Timestamp)
	}

	validateSensorField(t, "PM1_0", lastData.PM1_0, newData.PM1_0, 0.75)
	validateSensorField(t, "PM2_5", lastData.PM2_5, newData.PM2_5, 0.5)
	validateSensorField(t, "PM4_0", lastData.PM4_0, newData.PM4_0, 0.5)
	validateSensorField(t, "PM10_0", lastData.PM10_0, newData.PM10_0, 0.5)
	validateSensorField(t, "Temp", lastData.Temp, newData.Temp, 0.5)
	validateSensorField(t, "Hum", lastData.Hum, newData.Hum, 0.5)

}

func validateSensorField(t *testing.T, fieldName string, lastValue, newValue, tolerance float32) {
	if newValue < lastValue-tolerance || newValue > lastValue+tolerance {
		t.Errorf("%s o valor esperado deveria estar entre %.2f de %f. Recebido %f", fieldName, tolerance, lastValue, newValue)
	} else {
		fmt.Printf("%s: OK\n", fieldName)
	}
}
