package repositories

import (
	"cims/internal/core/domain"
	authPort "cims/internal/core/ports/auth"

	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) authPort.Repository {
	return &authRepository{db: db}
}

func (r authRepository) GetByUsername(username string) (*domain.Users, error) {

	var user domain.Users

	if err := r.db.Model(&user).Debug().
		Preload("Staff").
		Preload("Staff.Role").
		Preload("Staff.Campus").
		Where("username = ?", username).
		Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
