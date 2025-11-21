package answers

import "gorm.io/gorm"

type PGRepository struct {
	db *gorm.DB
}

func NewPGRepository(db *gorm.DB) *PGRepository {
	return &PGRepository{db: db}
}

func (r *PGRepository) Create(questionID uint, userID, text string) Answer {
	a := Answer{
		QuestionID: questionID,
		UserID:     userID,
		Text:       text,
	}
	r.db.Create(&a)
	return a
}

func (r *PGRepository) Get(id uint) (Answer, error) {
	var a Answer
	res := r.db.First(&a, id)
	if res.Error != nil {
		return Answer{}, res.Error
	}
	return a, nil
}

func (r *PGRepository) Delete(id uint) error {
	return r.db.Delete(&Answer{}, id).Error
}

func (r *PGRepository) GetByQuestion(questionID uint) []Answer {
	var list []Answer
	r.db.Where("question_id = ?", questionID).Find(&list)
	return list
}
