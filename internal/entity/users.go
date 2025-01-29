package entity

import "github.com/uptrace/bun"

type Users struct {
	bun.BaseModel `bun:"table:projects"`

	basicEntity
	FullName *string `json:"full_name" bun:"full_name"`
	UserName *string `json:"username" bun:"username"`
	Password *string `json:"password" bun:"password"`
	Role     *string `json:"role" bun:"role"`
	IsActive *bool   `json:"is_active" bun:"is_active"`
}
