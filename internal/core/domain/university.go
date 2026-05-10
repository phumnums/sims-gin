package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type University struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
	DeletedAt *gorm.DeletedAt
	DeletedBy string
}

type Campus struct {
	ID              uuid.UUID
	Name            string
	Status          string
	StatusUpdatedAt time.Time
	UniversityID    uuid.UUID
	CreatedAt       time.Time
	CreatedBy       string
	UpdatedAt       time.Time
	UpdatedBy       string
	DeletedAt       *gorm.DeletedAt
	DeletedBy       string
	University      University `gorm:"foreignkey:UniversityID;references:ID"`
}
