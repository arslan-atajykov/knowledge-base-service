package api

import (
	"encoding/json"
	"knowledge-base-service/internal/answers"
	"knowledge-base-service/internal/questions"
	"net/http"
	"strconv"
	"strings"
)

type QHandler struct {
	qService *questions.Service
	aService *answers.Service
}

func NewQHandler(q *questions.Service, a *answers.Service) *QHandler {
	return &QHandler{qService: q, aService: a}
}

func (h *QHandler) Route(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path == "/questions" || path == "/questions/" {
		switch r.Method {
		case http.MethodGet:
			h.GetAll(w, r)
		case http.MethodPost:
			h.Create(w, r)
		default:
			http.NotFound(w, r)
		}
		return
	}

	if strings.HasPrefix(path, "/questions/") && strings.HasSuffix(path, "/answers/") {
		idStr := strings.TrimSuffix(strings.TrimPrefix(path, "/questions/"), "/answers/")
		id, err := strconv.Atoi(strings.Trim(idStr, "/"))
		if err != nil {
			http.Error(w, "неверный id", http.StatusBadRequest)
			return
		}

		if r.Method == http.MethodPost {
			h.CreateAnswer(w, r, uint(id))
			return
		}

		http.NotFound(w, r)
		return
	}

	if strings.HasPrefix(path, "/questions/") {
		idStr := strings.TrimPrefix(path, "/questions/")
		id, err := strconv.Atoi(strings.Trim(idStr, "/"))
		if err != nil {
			http.Error(w, "неверный id", http.StatusBadRequest)
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
		return
	}

	http.NotFound(w, r)
}

func (h *QHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := h.qService.GetAll()
	json.NewEncoder(w).Encode(res)
}

func (h *QHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Text string `json:"text"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "неверный json", http.StatusBadRequest)
		return
	}

	q, err := h.qService.Create(body.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(q)
}

func (h *QHandler) GetByID(w http.ResponseWriter, r *http.Request, id uint) {
	q, err := h.qService.GetByID(id)
	if err != nil {
		http.Error(w, "не найдено", http.StatusNotFound)
		return
	}

	ans := h.aService.GetByQuestion(id)

	res := struct {
		questions.Question
		Answers []answers.Answer `json:"answers"`
	}{
		Question: q,
		Answers:  ans,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *QHandler) Delete(w http.ResponseWriter, r *http.Request, id uint) {

	ansList := h.aService.GetByQuestion(id)
	for _, a := range ansList {
		_ = h.aService.Delete(a.ID)
	}

	if err := h.qService.Delete(id); err != nil {
		http.Error(w, "не найдено", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"сообщение": "Вопрос и все ответы к этому вопросу удалены",
	})
}

func (h *QHandler) CreateAnswer(w http.ResponseWriter, r *http.Request, questionID uint) {

	if _, err := h.qService.GetByID(questionID); err != nil {
		http.Error(w, "вопрос не найден", http.StatusNotFound)
		return
	}

	var body struct {
		UserID string `json:"user_id"`
		Text   string `json:"text"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "ошибка json", http.StatusBadRequest)
		return
	}

	a, err := h.aService.Create(questionID, body.UserID, body.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(a)
}
