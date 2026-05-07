package repositories

import (
	"cims/internal/adapters/dto"
	"cims/internal/core/domain"
	studentPort "cims/internal/core/ports/student"
	"fmt"

	"gorm.io/gorm"
)

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) studentPort.Repository {
	return &studentRepository{db: db}
}

func (r *studentRepository) FindAll(params dto.SearchStudents) ([]domain.Students, int64, error) {

	var students []domain.Students
	var total int64

	db := r.db.Debug().Model(&students)
	page := params.Page
	size := params.Size

	// firstname
	if params.FirstName != "" {
		firstname := fmt.Sprintf("%%%s%%", params.FirstName)
		db = db.Where("first_name_th Like ? OR first_name_eng Like ?", firstname, firstname)
	}

	// lastname
	if params.LastName != "" {
		lastName := fmt.Sprintf("%%%s%%", params.LastName)
		db = db.Where("last_name_th Like ? OR last_name_eng LIKE ?", lastName, lastName)
	}

	// national id
	if params.NationalID != "" {
		db = db.Where("national_id = ?", params.NationalID)
	}

	// phone number
	if params.PhoneNumber != "" {
		db = db.Where("phone_number = ?", params.PhoneNumber)
	}

	// email
	if params.Email != "" {
		db = db.Where("email = ?", params.Email)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size

	if err := db.
		Where("campus_id = ?", params.CampusID).
		Limit(size).
		Offset(offset).
		Find(&students).Error; err != nil {
		return nil, 0, err
	}

	return students, total, nil
}
