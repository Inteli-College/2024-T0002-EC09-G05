package utils

import (
	"fmt"
	"time"
)

func PrintASCIIArt() {
	time.Sleep(5 * time.Second)
	fmt.Println(``)
	fmt.Println(``)
	fmt.Println(`    ___/\\\\\\\\\\\\____/\\\\\\\\\\\\\\\__       `)
	fmt.Println(`    __/\\\//////////____\/\\\///////////__       `)
	fmt.Println(`    __/\\\_______________\/\\\_____________      `)
	fmt.Println(`    __\/\\\____/\\\\\\\___\/\\\\\\\\\\\\_____    `)
	fmt.Println(`    ___\/\\\___\/////\\\___\////////////\\\___   `)
	fmt.Println(`     ___\/\\\_______\/\\\______________\//\\\__  `)
	fmt.Println(`      ___\/\\\_______\/\\\___/\\\________\/\\\__ `)
	fmt.Println(`       ___\//\\\\\\\\\\\\/___\//\\\\\\\\\\\\\/__ `)
	fmt.Println(`        ___\////////////______\/////////////___  `)
	fmt.Println(``)
	fmt.Println(`                      Grupo 5                    `)
	fmt.Println(`                All systems online!              `)
	fmt.Println(``)
	fmt.Println(`    Dashboard do InfluxDB: http://localhost:8086 `)
	fmt.Println(`    Dashboard do RabbitMQ: http://localhost:15672`)
	fmt.Println(`    Endpoint do servidor:  http://localhost:8080 `)
	fmt.Println(``)
	fmt.Println(``)
	
}
