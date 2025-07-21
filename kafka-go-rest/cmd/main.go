package main

import (
	"log"
	"net/http"

	"kafka-go-rest/router"
)

func main() {
	router := router.SetupRouter()
	log.Println("Starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
