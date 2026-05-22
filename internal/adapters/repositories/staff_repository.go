package repositories

import (
	"cims/internal/core/domain"
	staffPort "cims/internal/core/ports/staff"
	"context"

	"gorm.io/gorm"
)

type staffRepository struct {
	db *gorm.DB
}

func NewStaffRepository(db *gorm.DB) staffPort.Repository {
	return &staffRepository{db: db}
}

func (r *staffRepository) GetRoleByID(ctx context.Context, roleID string) (*domain.Roles, error) {
	var role domain.Roles

	if err := r.db.WithContext(ctx).
		Model(&role).
		Where("id = ?", roleID).
		Take(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *staffRepository) GetCampusByID(ctx context.Context, campusID string) (*domain.Campus, error) {
	var campus domain.Campus

	if err := r.db.WithContext(ctx).
		Model(&campus).
		Where("id = ?", campusID).
		Take(&campusID).Error; err != nil {
		return nil, err
	}

	return &campus, nil
}

func (r *staffRepository) CheckDuplicateUsername(ctx context.Context, username string) error {
	var user domain.Users

	if err := r.db.WithContext(ctx).
		Model(&user).
		Where("username = ?", username).
		Take(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *staffRepository) CheckDuplicatePhoneNumber(ctx context.Context, phoneNumber string) error {
	var staff domain.Staffs

	if err := r.db.WithContext(ctx).
		Model(&staff).
		Where("username = ?", phoneNumber).
		Take(&staff).Error; err != nil {
		return err
	}

	return nil
}

func (r *staffRepository) CheckDuplicateEmail(ctx context.Context, email string) error {
	var staff domain.Staffs

	if err := r.db.WithContext(ctx).
		Model(&staff).
		Where("username = ?", email).
		Take(&staff).Error; err != nil {
		return err
	}

	return nil
}

func (r *staffRepository) CreateUser(ctx context.Context, tx *gorm.DB, user domain.Users) error {

	if err := r.db.WithContext(ctx).
		Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *staffRepository) CreateStaff(ctx context.Context, tx *gorm.DB, staff domain.Staffs) error {

	if err := r.db.WithContext(ctx).
		Create(staff).Error; err != nil {
		return err
	}

	return nil
}
