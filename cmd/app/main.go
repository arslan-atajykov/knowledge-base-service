package main

import (
	"log"
	"net/http"

	"knowledge-base-service/internal/answers"
	"knowledge-base-service/internal/api"
	"knowledge-base-service/internal/questions"
)

func main() {

	qRepo := questions.NewRepository()
	aRepo := answers.NewRepository()

	qService := questions.NewService(qRepo)
	aService := answers.NewService(aRepo)

	router := api.NewRouter(qService, aService)

	host := ":8080"
	log.Printf("HTTP сервер запущен на %s", host)

	if err := http.ListenAndServe(host, router); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
