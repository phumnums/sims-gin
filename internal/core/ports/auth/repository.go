package auth

import (
	"cims/internal/core/domain"
)

type Repository interface {
	GetByUsername(username string) (*domain.Users, error)
}
