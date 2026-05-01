package domain

import (
	"time"
)

type Student struct {
	StudentID      string
	MiddleName     string
	FirstName      string
	LastName       string
	BirthDate      time.Time
	Gender         string
	NatinalID      string
	PassportID     string
	PhoneNumber    string
	Email          string
	StudentType    string
	Status         string
	StatusUpdateAt time.Time
	CampusID       string
}
