package answers

import "errors"

type RepositoryInterface interface {
	Create(questionID uint, userID, text string) Answer
	Get(id uint) (Answer, error)
	Delete(id uint) error
	GetByQuestion(questionID uint) []Answer
}

type Service struct {
	repo RepositoryInterface
}

func NewService(repo RepositoryInterface) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(questionID uint, userID, text string) (Answer, error) {
	if userID == "" {
		return Answer{}, errors.New("Нужен user_id")
	}
	if len(text) < 1 {
		return Answer{}, errors.New("Нужен text")
	}
	return s.repo.Create(questionID, userID, text), nil
}

func (s *Service) GetByID(id uint) (Answer, error) {
	return s.repo.Get(id)
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *Service) GetByQuestion(questionID uint) []Answer {
	return s.repo.GetByQuestion(questionID)
}
