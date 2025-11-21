package main

import (
	"log"
	"net/http"

	"knowledge-base-service/internal/answers"
	"knowledge-base-service/internal/api"
	"knowledge-base-service/internal/db"
	"knowledge-base-service/internal/questions"
)

func main() {

	dbConn, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	dbConn.AutoMigrate(&questions.Question{}, &answers.Answer{})

	qRepo := questions.NewPGRepository(dbConn)
	aRepo := answers.NewPGRepository(dbConn)

	qService := questions.NewService(qRepo)
	aService := answers.NewService(aRepo)

	router := api.NewRouter(qService, aService)

	host := ":8080"
	log.Printf("HTTP сервер запущен на %s", host)

	if err := http.ListenAndServe(host, router); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
