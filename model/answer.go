package model

import "time"

type Answer struct {
	ID         int       `json:"id"`
	QuestionID int       `json:"question_id"`
	Content    string    `json:"content"`
	Username   string    `json:"username"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
