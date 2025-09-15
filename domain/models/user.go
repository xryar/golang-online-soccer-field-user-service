package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uint
	UUID        uuid.UUID
	Name        string
	Pass        string
	PhoneNumber string
	Email       string
	RoleID      uint
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	Role        Role `gorm:"foreignKey:role_id;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
