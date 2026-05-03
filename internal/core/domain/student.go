package domain

import (
	"time"
)

type Students struct {
	ID             string
	FirstName      string
	MiddleName     string
	LastName       string
	BirthDate      time.Time
	Gender         string
	NationalID     string
	PassportID     string
	PhoneNumber    string
	Email          string
	StudentType    string
	Status         string
	InactiveReason string
	StatusUpdateAt time.Time
	CampusID       string
	CreatedAt      time.Time
	CreatedBy      string
	UpdatedAt      time.Time
	UpdatedBy      string
	DeletedAt      time.Time
	DeletedBy      string
	Campus         Campus `gorm:"foreignkey:CampusID;references:ID"`
}
