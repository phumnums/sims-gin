package student

import (
	"cims/internal/adapters/dto"
)

type Service interface {
	GetAll(param dto.SearchStudents) ([]dto.StudentResponse, int, int, int64, int, error)
}
