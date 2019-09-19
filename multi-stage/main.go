package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alwaysbespoke/docker/multi-stage/handlers"
)

func main() {
	// start server
	port := "8080"
	p, ok := os.LookupEnv("PORT")
	if ok {
		port = p
	}
	address := ":" + port
	fmt.Println("Starting server: ", address)
	http.HandleFunc("/", handlers.EchoHandler)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal(err)
	}
}
