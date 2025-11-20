package answers

import "time"

type Answer struct {
	ID         uint      `json:"id"`
	QuestionID uint      `json:"question_id"`
	UserID     string    `json:"user_id"`
	Text       string    `json:"text"`
	CreatedAt  time.Time `json:"created_at"`
}
