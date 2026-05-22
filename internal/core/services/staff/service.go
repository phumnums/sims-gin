package user

import (
	"cims/internal/adapters/dto"
	"cims/internal/adapters/middleware"
	"cims/internal/core/domain"
	staffPort "cims/internal/core/ports/staff"
	"cims/pkg/utils/password"
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type service struct {
	repo staffPort.Repository
	db   *gorm.DB
}

func NewStaffService(repo staffPort.Repository, db *gorm.DB) staffPort.Service {
	return &service{
		repo: repo,
		db:   db,
	}
}

func (s *service) Create(ctx context.Context, auth middleware.AuthContext, req dto.CreateStaffRequest) (int, error) {

	// check role id
	role, err := s.repo.GetRoleByID(ctx, req.RoleID)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return http.StatusInternalServerError, errors.New("server error")
		} else {
			return http.StatusNotFound, errors.New("role not found")
		}
	}

	// check campus id
	campus, err := s.repo.GetCampusByID(ctx, req.CampusID)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return http.StatusInternalServerError, errors.New("server error")
		} else {
			return http.StatusNotFound, errors.New("campus not found")
		}
	}

	// check duplicate username
	if err := s.repo.CheckDuplicateUsername(ctx, req.Username); err != nil {
		if err != gorm.ErrRecordNotFound {
			return http.StatusInternalServerError, errors.New("server error")
		}
	} else {
		return http.StatusConflict, errors.New("username is duplicate")
	}

	// check phone number username
	if err := s.repo.CheckDuplicatePhoneNumber(ctx, req.PhoneNumber); err != nil {
		if err != gorm.ErrRecordNotFound {
			return http.StatusInternalServerError, errors.New("server error")
		}
	} else {
		return http.StatusConflict, errors.New("phone number is duplicate")
	}

	// check email username
	if err := s.repo.CheckDuplicateEmail(ctx, req.Email); err != nil {
		if err != gorm.ErrRecordNotFound {
			return http.StatusInternalServerError, errors.New("server error")
		}
	} else {
		return http.StatusConflict, errors.New("email is duplicate")
	}

	passwordHash, err := password.HashPassword(req.Password)
	if err != nil {
		return http.StatusInternalServerError, errors.New("server error")
	}

	tx := s.db.WithContext(ctx)
	staffID := uuid.New()
	// created user
	userData := domain.Users{
		ID:           uuid.New(),
		Username:     req.Username,
		PasswordHash: passwordHash,
		StaffID:      staffID,
		CreatedAt:    time.Now(),
		CreatedBy:    auth.UserID,
		UpdatedAt:    time.Now(),
		UpdatedBy:    auth.UserID,
	}

	staffData := domain.Staffs{
		ID:              staffID,
		FirstNameTH:     req.FirstNameTH,
		MiddleNameTH:    &req.MiddleNameTH,
		LastNameTH:      req.LastNameTH,
		FirstNameENG:    req.FirstNameENG,
		MiddleNameENG:   &req.LastNameENG,
		LastNameENG:     req.LastNameENG,
		PhoneNumber:     req.PhoneNumber,
		Email:           req.Email,
		Status:          "Active",
		StatusUpdatedAt: time.Now(),
		RoleID:          role.ID,
		CampusID:        campus.ID,
		CreatedAt:       time.Now(),
		CreatedBy:       auth.UserID,
		UpdatedAt:       time.Now(),
		UpdatedBy:       auth.UserID,
	}

	tx.Begin()
	if err := s.repo.CreateUser(ctx, tx, userData); err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, errors.New("server error")
	}
	if err := s.repo.CreateStaff(ctx, tx, staffData); err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, errors.New("server error")
	}
	tx.Commit()

	return http.StatusOK, nil
}
