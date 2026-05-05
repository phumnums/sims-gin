package domain

import (
	"time"
)

type Students struct {
	ID              string
	FirstNameTH     string
	MiddleNameTH    string
	LastNameTH      string
	FirstNameENG    string
	MiddleNameENG   string
	LastNameENG     string
	BirthDate       time.Time
	Gender          string
	NationalID      string
	PassportID      string
	PhoneNumber     string
	Email           string
	StudentType     string
	Status          string
	InactiveReason  string
	StatusUpdatedAt time.Time
	CampusID        string
	CreatedAt       time.Time
	CreatedBy       string
	UpdatedAt       time.Time
	UpdatedBy       string
	DeletedAt       time.Time
	DeletedBy       string
	Campus          Campus `gorm:"foreignkey:CampusID;references:ID"`
}
