package student

import "cims/internal/core/domain"

type Service interface {
	SearchStudents(
		firstName string,
		lastName string,
		nationalID string,
		phoneNumber string,
		email string,
		campusID string,
	) ([]domain.Students, error)
}
