package entity

import "github.com/uptrace/bun"

type Answers struct {
	bun.BaseModel `bun:"table:projects"`

	basicEntity
	QuestionId  *int    `json:"question_id" bun:"question_id"`
	AnswerText  *string `json:"answer_text" bun:"answer_text"`
	IsCorrect   *bool   `json:"is_correct" bun:"is_correct"`
	OrderNumber *int    `json:"order_number" bun:"order_number"`
}
