package controllers

import (
	"encoding/json"
	"kafka-go-rest/models"
	"kafka-go-rest/producer"
	"net/http"
)

func ProduceHandler(w http.ResponseWriter, r *http.Request) {
	var msg models.Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err = producer.ProduceMessage(msg.Value)
	if err != nil {
		http.Error(w, "Failed to produce", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message sent to Kafka"))
}
