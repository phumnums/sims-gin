package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Roles struct {
	ID              uuid.UUID
	Name            string
	Status          string
	StatusUpdatedAt time.Time
	CreatedAt       time.Time
	CreatedBy       string
	UpdatedAt       time.Time
	UpdatedBy       string
	DeletedAt       *gorm.DeletedAt
	DeletedBy       *string
}
