package questions

import (
	"gorm.io/gorm"
)

type PGRepository struct {
	db *gorm.DB
}

func NewPGRepository(db *gorm.DB) *PGRepository {
	return &PGRepository{db: db}
}

func (r *PGRepository) Create(text string) Question {
	q := Question{Text: text}
	r.db.Create(&q)
	return q
}

func (r *PGRepository) GetAll() []Question {
	var qs []Question
	r.db.Find(&qs)
	return qs
}

func (r *PGRepository) Get(id uint) (Question, error) {
	var q Question
	res := r.db.First(&q, id)
	if res.Error != nil {
		return Question{}, res.Error
	}
	return q, nil
}

func (r *PGRepository) Delete(id uint) error {
	return r.db.Delete(&Question{}, id).Error
}
