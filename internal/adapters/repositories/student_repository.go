package repositories

import (
	"cims/internal/core/domain"
	studentPort "cims/internal/core/ports/student"

	"gorm.io/gorm"
)

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) studentPort.Repository {
	return &studentRepository{db: db}
}

func (r *studentRepository) Search(
	firstName string,
	lastName string,
	nationalID string,
	phoneNumber string,
	email string,
	campusID string,
) ([]domain.Students, error) {
	var students []domain.Students

	err := r.db.Where("campus_id = ?", campusID).Find(students).Error

	if err != nil {
		return nil, err
	}

	return students, nil
}
