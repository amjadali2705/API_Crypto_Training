package router

import (
	"kafka-go-rest/consumer"
	"kafka-go-rest/controllers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/produce", controllers.ProduceHandler).Methods("POST")
	consumer.StartConsumer()
	return r
}
