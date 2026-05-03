package student

import (
	"cims/internal/core/domain"
)

type Repository interface {
	Search(
		firstName string,
		lastName string,
		nationalID string,
		phoneNumber string,
		email string,
		campusID string,
	) ([]domain.Students, error)
}
