package answers

import (
	"errors"
	"sync"
	"time"
)

type Repository struct {
	mu      sync.RWMutex
	answers map[uint]Answer
	nextID  uint
}

func NewRepository() *Repository {
	return &Repository{
		answers: make(map[uint]Answer),
		nextID:  1,
	}
}

func (r *Repository) Create(questionID uint, userID, text string) Answer {
	r.mu.Lock()
	defer r.mu.Unlock()
	a := Answer{
		ID:         r.nextID,
		QuestionID: questionID,
		UserID:     userID,
		Text:       text,
		CreatedAt:  time.Now(),
	}
	r.answers[r.nextID] = a
	r.nextID++
	return a
}

func (r *Repository) Get(id uint) (Answer, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	a, ok := r.answers[id]
	if !ok {
		return Answer{}, errors.New("Не найдено")
	}
	return a, nil
}

func (r *Repository) Delete(id uint) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.answers[id]; !ok {
		return errors.New("Не найдено")
	}
	delete(r.answers, id)
	return nil
}

func (r *Repository) GetByQuestion(questionID uint) []Answer {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := []Answer{}
	for _, a := range r.answers {
		if a.QuestionID == questionID {
			result = append(result, a)
		}
	}
	return result
}
