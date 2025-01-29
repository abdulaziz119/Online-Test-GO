package entity

import "github.com/uptrace/bun"

type TestAttempts struct {
	bun.BaseModel `bun:"table:projects"`

	basicEntity
	UserId     *int    `json:"user_id" bun:"user_id"`
	TestId     *int    `json:"test_id" bun:"test_id"`
	StartedAt  *string `json:"started_at" bun:"started_at"`
	FinishedAt *string `json:"finished_at" bun:"finished_at"`
	Score      *int    `json:"score" bun:"score"`
	Status     *string `json:"status" bun:"status"`
}
