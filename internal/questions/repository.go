package questions

import (
	"errors"
	"sync"
	"time"
)

type Repository struct {
	mu        sync.RWMutex
	questions map[uint]Question
	nextID    uint
}

func NewRepository() *Repository {
	return &Repository{
		questions: make(map[uint]Question),
		nextID:    1,
	}
}

func (r *Repository) Create(text string) Question {
	r.mu.Lock()
	defer r.mu.Unlock()
	q := Question{
		ID:        r.nextID,
		Text:      text,
		CreatedAt: time.Now(),
	}
	r.questions[r.nextID] = q
	r.nextID++
	return q
}

func (r *Repository) GetAll() []Question {
	r.mu.RLock()
	defer r.mu.RUnlock()

	results := make([]Question, 0, len(r.questions))

	for _, q := range r.questions {
		results = append(results, q)
	}
	return results
}

func (r *Repository) Get(id uint) (Question, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	q, ok := r.questions[id]
	if !ok {
		return Question{}, errors.New("не найдено")
	}
	return q, nil
}

func (r *Repository) Delete(id uint) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.questions[id]; !ok {
		return errors.New("Не найдено")
	}
	delete(r.questions, id)
	return nil
}
