package student

import (
	"cims/internal/adapters/dto"
	"cims/internal/adapters/mappers"
	studentPort "cims/internal/core/ports/student"
	"errors"
	"net/http"
)

type service struct {
	repo studentPort.Repository
}

func NewService(repo studentPort.Repository) studentPort.Service {
	return &service{repo: repo}
}

func (s *service) GetAll(params dto.SearchStudents) ([]dto.StudentResponse, int, int, int64, int, error) {

	if params.Page < 1 {
		params.Page = 1
	}
	if params.Size < 1 {
		params.Size = 10
	}

	students, total, err := s.repo.FindAll(params)
	if err != nil {
		return nil, 0, 0, 0, http.StatusInternalServerError, errors.New("server error")
	}

	studentMapper := mappers.StudentResponseMapper(students)

	return studentMapper, params.Page, params.Size, total, http.StatusOK, nil
}
