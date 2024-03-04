package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Sensor struct {
	nameType        string
	delta           float64
	min             float64
	max             float64
	connection      bool
	counter         int
	status          string
	startWave       bool
	currentPosition float64
	goPosition      float64
	waveSize        int
}

func NewSensor(name string, min float64, max float64) *Sensor {
	return &Sensor{
		nameType:        name,
		min:             min,
		max:             max,
		connection:      false,
		counter:         10,
		status:          "new wave",
		startWave:       true,
		currentPosition: 0,
		goPosition:      0,
		waveSize:        0,
	}
}

func (s *Sensor) On(broker map[string]string, test map[string]interface{}) {
	// Connect to the broker
	s.connection = true
	s.counter = test["sec"].(int)
	for s.connected(test) {
		message := s.generateWaveData()
		fmt.Println(message)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Publicação encerrada")
}

func (s *Sensor) Off() {
	s.connection = false
	fmt.Println("Publicação encerrada")
}

func (s *Sensor) connected(test map[string]interface{}) bool {
	if test["tested"].(bool) == true {
		s.counter -= 1
		if s.counter < 0 {
			s.Off()
		}
	}
	return s.connection
}

func (s *Sensor) generateWaveData() string {
	var value float64

	if s.status == "new wave" {
		s.waveSize = 5
		if s.startWave == true {
			s.currentPosition = float64(rand.Intn(int(s.max-s.min)) + int(s.min))
			s.startWave = false
			value = s.currentPosition
		}
		s.status = "wave"
	} else if s.status == "wave" {
		s.waveSize -= 1
		value = s.currentPosition + math.Round(math.Abs(rand.NormFloat64()*5)*math.Cos(s.delta)*10)/10
		if s.waveSize <= 0 {
			s.status = "transtion"
			s.goPosition = float64(rand.Intn(int(s.max-s.min)) + int(s.min))
		}
	} else if s.status == "transtion" {
		fmt.Printf("Indo de %.2f para %.2f\n", s.currentPosition, s.goPosition)
		posOrNeg := 0.0
		difference := s.currentPosition - s.goPosition

		if difference > 0 {
			posOrNeg = 1.0
		} else {
			posOrNeg = -1.0
		}

		fmt.Println(difference, posOrNeg)
		s.currentPosition -= posOrNeg
		value = s.currentPosition
		if s.currentPosition == s.goPosition {
			s.currentPosition = s.goPosition
			s.status = "new wave"
		}
	} else {
		fmt.Println("Erro")
	}

	if value > s.max {
		value = s.max
	}
	if value < s.min {
		value = s.min
	}

	return fmt.Sprintf("Valores do Sensor Simulado: %.2f", value)
}

func main() {
	s := NewSensor("sensor", -10, 10)
	broker := map[string]string{"link": "2638385848004a349ca166f397873de7.s1.eu.hivemq.cloud", "port": "8883"}
	test := map[string]interface{}{"tested": false, "sec": 0}

	s.On(broker, test)
}
