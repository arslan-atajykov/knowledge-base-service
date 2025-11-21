package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"knowledge-base-service/internal/answers"
	"knowledge-base-service/internal/api"
	"knowledge-base-service/internal/questions"
)

func TestCreateQuestion(t *testing.T) {
	// In memory проверка
	qRepo := questions.NewRepository()
	aRepo := answers.NewRepository()

	qService := questions.NewService(qRepo)
	aService := answers.NewService(aRepo)

	server := httptest.NewServer(api.NewRouter(qService, aService))
	defer server.Close()

	body := []byte(`{"text": "Столица Англии?"}`)

	resp, err := http.Post(server.URL+"/questions", "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("ошибка при загрузке вопроас /questions: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Ожидаемый статус 201, получили %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("неверный JSON: %v", err)
	}

	if result["id"] == nil {
		t.Fatalf("ожидали ID")
	}

	if result["text"] != "Столица Англии?" {
		t.Fatalf("текст не соответствует, got: %v", result["text"])
	}
}
