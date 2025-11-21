package api

import (
	"knowledge-base-service/internal/answers"
	"knowledge-base-service/internal/questions"
	"net/http"
)

func NewRouter(qService *questions.Service,
	aService *answers.Service) http.Handler {
	mux := http.NewServeMux()

	qHandler := NewQHandler(qService, aService)
	aHandler := NewAHandler(aService)

	mux.HandleFunc("/questions/", qHandler.Route)
	mux.HandleFunc("/questions", qHandler.Route)

	mux.HandleFunc("/answers/", aHandler.Route)
	mux.HandleFunc("/answers", aHandler.Route)

	// mux.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusOK)
	// 	_, _ = w.Write([]byte("ok"))
	// })
	return mux
}
