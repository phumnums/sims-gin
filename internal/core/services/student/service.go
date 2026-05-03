package student

import (
	"cims/internal/core/domain"
	studentPort "cims/internal/core/ports/student"
	"errors"
)

type service struct {
	repo studentPort.Repository
}

func NewService(repo studentPort.Repository) studentPort.Service {
	return &service{repo: repo}
}

func (s *service) SearchStudents(
	firstName string,
	lastName string,
	nationalID string,
	phoneNumber string,
	email string,
	campusID string,
) ([]domain.Students, error) {

	if campusID == "" {
		return nil, errors.New("")
	}

	students, err := s.repo.Search(firstName, lastName, nationalID, phoneNumber, email, campusID)
	if err != nil {
		return nil, errors.New("")
	}

	return students, nil
}
