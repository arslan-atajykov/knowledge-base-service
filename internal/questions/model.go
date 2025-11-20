package questions

import "time"

type Question struct {
	ID        uint      `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
