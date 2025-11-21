package questions

import "errors"

type QuestionRepo interface {
	Create(text string) Question
	GetAll() []Question
	Get(id uint) (Question, error)
	Delete(id uint) error
}
type Service struct {
	repo QuestionRepo
}

func NewService(r QuestionRepo) *Service {
	return &Service{repo: r}
}

func (s *Service) Create(text string) (Question, error) {
	if len(text) < 1 {
		return Question{}, errors.New("текст очень короткий")
	}
	return s.repo.Create(text), nil
}

func (s *Service) GetAll() []Question {
	return s.repo.GetAll()
}

func (s *Service) GetByID(id uint) (Question, error) {
	return s.repo.Get(id)
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}
