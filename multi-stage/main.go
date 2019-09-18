package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alwaysbespoke/docker/multi-stage/handlers"
)

func main() {
	// start server
	address := ":8080"
	fmt.Println("Starting server: ", address)
	http.HandleFunc("/", handlers.EchoHandler)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal(err)
	}
}
