package utils

import (
	"fmt"
	"time"
)

func PrintASCIIArt() {
	time.Sleep(5 * time.Second)
	fmt.Println(`---------------------------------------------------------------`)
	fmt.Println(`|                                                             |`)
	fmt.Println(`|    ██████╗ ██████╗  ██████╗   ███████╗███████╗███████╗      |`)
	fmt.Println(`|    ██╔══██╗██╔══██╗██╔═══██╗  ██╔════╝██╔════╝██╔════╝      |`)
	fmt.Println(`|    ██████╔╝██████╔╝██║   ██║  ███████╗█████╗  █████╗        |`)
	fmt.Println(`|    ██╔═══╝ ██╔══██╗██║   ██║  ╚════██║██╔══╝  ██╔══╝        |`)
	fmt.Println(`|    ██║     ██║  ██║╚██████╔╝  ███████║███████╗███████╗██╗   |`)
	fmt.Println(`|    ╚═╝     ╚═╝  ╚═╝ ╚═════╝   ╚══════╝╚══════╝╚══════╝╚═╝   |`)													
	fmt.Println(`|                                                             |`)
	fmt.Println(`|                          Grupo 5                            |`)
	fmt.Println(`|                    All systems online!                      |`)
	fmt.Println(`|                                                             |`)
	fmt.Println(`|       Dashboard do InfluxDB: http://localhost:8086          |`)
	fmt.Println(`|       Dashboard do RabbitMQ: http://localhost:15672         |`)
	fmt.Println(`|       Endpoint do servidor:  http://localhost:8080          |`)
	fmt.Println(`|                                                             |`)
	fmt.Println(`---------------------------------------------------------------`)

	
}
