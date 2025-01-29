package entity

import "github.com/uptrace/bun"

type Questions struct {
	bun.BaseModel `bun:"table:projects"`

	basicEntity
	TestId       *int    `json:"test_id" bun:"test_id"`
	QuestionText *string `json:"question_text" bun:"question_text"`
	Points       *int    `json:"points" bun:"points"`
	QuestionType *string `json:"question_type" bun:"question_type"`
	OrderNumber  *int    `json:"order_number" bun:"order_number"`
}
