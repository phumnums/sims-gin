package student

import (
	"cims/internal/adapters/dto"
)

type Service interface {
	SearchStudents(param dto.SearchStudents) ([]dto.StudentResponse, int, int, int64, error)
}
