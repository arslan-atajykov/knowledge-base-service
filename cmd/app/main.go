package main

import (
	"knowledge-base-service/internal/api"
	"log"
	"net/http"
)

func main() {
	router := api.Router()

	host_add := ":8080"
	log.Printf("включение HTTP сервера на %s\n", host_add)

	if err := http.ListenAndServe(host_add, router); err != nil {
		log.Fatalf("сервер не запустился: %v", err)
	}
}
