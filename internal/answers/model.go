package answers

import "time"

type Answer struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	QuestionID uint      `json:"question_id" gorm:"index;not null"`
	UserID     string    `json:"user_id" gorm:"type:varchar(255);not null"`
	Text       string    `json:"text" gorm:"type:text;not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
}
