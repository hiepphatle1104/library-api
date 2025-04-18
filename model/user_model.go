package model

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID       uuid.UUID `bun:",pk,type:uuid" json:"-"`
	FullName string    `bun:"full_name,notnull" json:"fullname"`
	UserName string    `bun:"user_name,notnull" json:"username"`
	Password string    `bun:"password,notnull" json:"password"`
}
