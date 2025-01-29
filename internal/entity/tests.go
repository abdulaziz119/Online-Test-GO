package entity

import "github.com/uptrace/bun"

type Test struct {
	bun.BaseModel `bun:"table:projects"`

	basicEntity
	SubjectId       *int    `json:"subject_id" bun:"subject_id"`
	TeacherId       *int    `json:"teacher_id" bun:"teacher_id"`
	Title           *string `json:"title" bun:"title"`
	Description     *string `json:"description" bun:"description"`
	DurationMinutes *int    `json:"duration_minutes" bun:"duration_minutes"`
	PassPercent     *int    `json:"pass_percent" bun:"pass_percent"`
	StartTime       *string `json:"start_time" bun:"start_time"`
	EndTime         *string `json:"end_time" bun:"end_time"`
	IsActive        *bool   `json:"is_active" bun:"is_active"`
}
