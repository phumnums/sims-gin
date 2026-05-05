package mappers

import (
	"cims/internal/adapters/dto"
	"cims/internal/core/domain"
)

func StudentResponseMapper(students []domain.Students) []dto.StudentResponse {
	var result []dto.StudentResponse
	for _, s := range students {
		data := dto.StudentResponse{
			ID:              s.ID,
			FirstNameTH:     s.FirstNameTH,
			MiddleNameTH:    s.MiddleNameTH,
			LastNameTH:      s.LastNameTH,
			FirstNameENG:    s.FirstNameENG,
			MiddleNameENG:   s.MiddleNameENG,
			LastNameENG:     s.LastNameENG,
			BirthDate:       s.BirthDate.Format("2006-01-02"),
			Gender:          s.Gender,
			NationalID:      s.NationalID,
			PassportID:      s.PassportID,
			PhoneNumber:     s.PhoneNumber,
			Email:           s.Email,
			StudentType:     s.StudentType,
			Status:          s.Status,
			InactiveReason:  s.InactiveReason,
			StatusUpdatedAt: s.StatusUpdatedAt.Format("2006-01-02"),
			CreatedAt:       s.CreatedAt.Format("2006-01-02"),
			CreatedBy:       s.CreatedBy,
			UpdatedAt:       s.UpdatedAt.Format("2006-01-02"),
			UpdatedBy:       s.UpdatedBy,
		}
		result = append(result, data)
	}

	return result
}
