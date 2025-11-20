package api

import (
	"net/http"
)

func Router() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	return mux
}
