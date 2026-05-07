package student

import (
	"cims/internal/adapters/dto"
	"cims/internal/core/domain"
)

type Repository interface {
	FindAll(params dto.SearchStudents) ([]domain.Students, int64, error)
}
