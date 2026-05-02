package domain

import (
	"time"

	"github.com/google/uuid"
)

type Roles struct {
	ID              uuid.UUID
	name            string
	Status          string
	StatusUpdatedAt time.Time
	CreatedAt       time.Time
	CreatedBy       string
	UpdatedAt       time.Time
	UpdatedBy       string
	DeletedAt       time.Time
	DeletedBy       string
}
