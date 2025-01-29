package entity

import "github.com/uptrace/bun"

type Subjects struct {
	bun.BaseModel `bun:"table:projects"`

	basicEntity
	Name        *string `json:"name" bun:"name"`
	Description *string `json:"description" bun:"description"`
	IsActive    *bool   `json:"is_active" bun:"is_active"`
}
