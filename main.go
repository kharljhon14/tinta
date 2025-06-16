package main

import (
	"log"

	"github.com/kharljhon14/tinta/cmd/api"
)

func main() {

	server := api.NewServer()

	err := server.Start(":8080")
	if err != nil {
		log.Fatal("unable to start server:", err.Error())
	}
}
