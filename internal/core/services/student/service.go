package student

import (
	"cims/internal/adapters/dto"
	"cims/internal/adapters/mappers"
	studentPort "cims/internal/core/ports/student"
)

type service struct {
	repo studentPort.Repository
}

func NewService(repo studentPort.Repository) studentPort.Service {
	return &service{repo: repo}
}

func (s *service) SearchStudents(params dto.SearchStudents) ([]dto.StudentResponse, int, int, int64, error) {

	if params.Page < 1 {
		params.Page = 1
	}
	if params.Size < 1 {
		params.Size = 10
	}

	students, total, err := s.repo.Search(params)
	if err != nil {
		return nil, 0, 0, 0, err
	}

	studentMapper := mappers.StudentResponseMapper(students)

	return studentMapper, params.Page, params.Size, total, nil
}
