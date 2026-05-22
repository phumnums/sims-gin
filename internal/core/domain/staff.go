package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Staffs struct {
	ID              uuid.UUID
	FirstNameTH     string
	MiddleNameTH    *string
	LastNameTH      string
	FirstNameENG    string
	MiddleNameENG   *string
	LastNameENG     string
	PhoneNumber     string
	Email           string
	Status          string
	StatusUpdatedAt time.Time
	InactiveReason  *string
	RoleID          uuid.UUID
	CampusID        uuid.UUID
	CreatedAt       time.Time
	CreatedBy       string
	UpdatedAt       time.Time
	UpdatedBy       string
	DeletedAt       *gorm.DeletedAt
	DeletedBy       *string
	Role            Roles  `gorm:"foreignkey:RoleID;references:ID"`
	Campus          Campus `gorm:"foreignkey:CampusID;references:ID"`
}

type Users struct {
	ID           uuid.UUID
	Username     string
	PasswordHash string
	StaffID      uuid.UUID
	CreatedAt    time.Time
	CreatedBy    string
	UpdatedAt    time.Time
	UpdatedBy    string
	DeletedAt    *gorm.DeletedAt
	DeletedBy    *string
	Staff        Staffs `gorm:"foreignkey:StaffID;references:ID"`
}
