package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	config := EnvConfig()
	router := NewRouter()

	fmt.Printf("Starting server on port %s...\n", config.Port)
	log.Fatal(http.ListenAndServe(config.AddressString(), router))
}
