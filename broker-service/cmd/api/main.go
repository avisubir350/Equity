package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {
}

func main() {
	log.Printf("Starting broker service on port : %s", webPort)

	app := Config{}

	log.Print("Making call to http server")

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	log.Print("Now it will listen and serve")
	err := server.ListenAndServe()
	if err != nil {
		log.Print("Error occured", err)
	}

	log.Printf("Broker service is started")
}
