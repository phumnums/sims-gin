package auth

import (
	"cims/internal/adapters/dto"
	authPort "cims/internal/core/ports/auth"
	"cims/pkg/utils/jwt"
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo authPort.Repository
}

func NewAuthService(repo authPort.Repository) authPort.Service {
	return &service{repo: repo}
}

func (s *service) Login(req dto.LoginRequest) (string, int, error) {

	// TODO:

	// check user form DB
	user, err := s.repo.GetByUsername(req.Username)
	fmt.Println("user", user)
	if err != nil {
		return "", http.StatusNotFound, errors.New("user not found")
	}

	// compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return "", http.StatusBadRequest, errors.New("password is incorrect")
	}

	// token, err := jwt.GenerageToken("staff-id", "admin", "campus_id")
	token, err := jwt.GenerageToken(user.StaffID.String(), user.Staff.Role.Name, user.Staff.CampusID.String())
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	return token, http.StatusOK, nil
}
