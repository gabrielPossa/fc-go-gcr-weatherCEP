package main

import (
	"fmt"

	"github.com/gabrielPossa/fc-go-gcr-weatherCEP/internal"
	"github.com/gabrielPossa/fc-go-gcr-weatherCEP/pkg/webserver"
)

func main() {
	httpServer := webserver.NewWebServer(":8080")
	httpServer.AddHandler("/cep/{CEP}/weather", webserver.GET, internal.GetWeatherByCEP)
	fmt.Println("Starting web server on port", "8080")
	httpServer.Start()
}
