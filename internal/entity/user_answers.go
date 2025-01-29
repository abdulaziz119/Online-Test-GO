package entity

import "github.com/uptrace/bun"

type UserAnswers struct {
	bun.BaseModel `bun:"table:projects"`

	basicEntity
	AttemptId  *int    `json:"attempt_id" bun:"attempt_id"`
	QuestionId *int    `json:"question_id" bun:"question_id"`
	AnswerId   *int    `json:"answer_id" bun:"answer_id"`
	AnsweredAt *string `json:"answered_at" bun:"answered_at"`
}
