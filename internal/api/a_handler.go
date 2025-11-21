package api

import (
	"encoding/json"
	"knowledge-base-service/internal/answers"
	"net/http"
	"strconv"
	"strings"
)

type AHandler struct {
	aService *answers.Service
}

func NewAHandler(a *answers.Service) *AHandler {
	return &AHandler{aService: a}
}

func (h *AHandler) Route(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if !strings.HasPrefix(path, "/answers/") {
		http.NotFound(w, r)
		return
	}

	idStr := strings.TrimPrefix(path, "/answers/")
	id, err := strconv.Atoi(strings.Trim(idStr, "/"))
	if err != nil {
		http.Error(w, "Неправильный id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.GetByID(w, r, uint(id))
	case http.MethodDelete:
		h.Delete(w, r, uint(id))
	default:
		http.NotFound(w, r)
	}
}

func (h *AHandler) GetByID(w http.ResponseWriter, r *http.Request, id uint) {
	a, err := h.aService.GetByID(id)
	if err != nil {
		http.Error(w, "не найден ответ", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(a)
}

func (h *AHandler) Delete(w http.ResponseWriter, r *http.Request, id uint) {
	if err := h.aService.Delete(id); err != nil {
		http.Error(w, "не найден ответ", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
