package utils

import (
	"fmt"
	"time"
)

func PrintASCIIArt() {
	time.Sleep(5 * time.Second)
	fmt.Println(``)
	fmt.Println(` ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓`)
	fmt.Println(` ┃                                                             ┃`)
	fmt.Println(` ┃                        Grupo 5's                            ┃`)
	fmt.Println(` ┃                                                             ┃`)
	fmt.Println(` ┃    ██████╗ ██████╗  ██████╗   ███████╗███████╗███████╗      ┃`)
	fmt.Println(` ┃    ██╔══██╗██╔══██╗██╔═══██╗  ██╔════╝██╔════╝██╔════╝      ┃`)
	fmt.Println(` ┃    ██████╔╝██████╔╝██║   ██║  ███████╗█████╗  █████╗        ┃`)
	fmt.Println(` ┃    ██╔═══╝ ██╔══██╗██║   ██║  ╚════██║██╔══╝  ██╔══╝        ┃`)
	fmt.Println(` ┃    ██║     ██║  ██║╚██████╔╝  ███████║███████╗███████╗██╗   ┃`)
	fmt.Println(` ┃    ╚═╝     ╚═╝  ╚═╝ ╚═════╝   ╚══════╝╚══════╝╚══════╝╚═╝   ┃`)
	fmt.Println(` ┃                                                             ┃`)
	fmt.Println(` ┃    https://github.com/Inteli-College/2024-T0002-EC09-G05    ┃`)
	fmt.Println(` ┃                                                             ┃`)
	fmt.Println(` ┃                    All systems online!                      ┃`)
	fmt.Println(` ┃                                                             ┃`)
	fmt.Println(` ┃       Dashboard do InfluxDB: http://localhost:8086          ┃`)
	fmt.Println(` ┃       Dashboard do RabbitMQ: http://localhost:15672         ┃`)
	fmt.Println(` ┃       Endpoint do servidor:  http://localhost:8080          ┃`)
	fmt.Println(` ┃                                                             ┃`)
	fmt.Println(` ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛`)
	fmt.Println(``)

}
