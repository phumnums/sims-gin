package domain

import (
	"time"

	"github.com/google/uuid"
)

type Staffs struct {
	ID             uuid.UUID
	FirstName      string
	MiddleName     string
	LastName       string
	PhoneNumber    string
	Email          string
	Status         string
	StatusUpdateAt time.Time
	InactiveReason string
	RoleID         uuid.UUID
	CampusID       uuid.UUID
	CreatedAt      time.Time
	CreatedBy      string
	UpdatedAt      time.Time
	UpdatedBy      string
	DeletedAt      time.Time
	DeletedBy      string
	Role           Roles  `gorm:"foreignkey:RoleID;references:ID"`
	Campus         Campus `gorm:"foreignkey:CampusID;references:ID"`
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
	DeletedAt    time.Time
	DeletedBy    string
	Staff        Staffs `gorm:"foreignkey:StaffID;references:ID"`
}
