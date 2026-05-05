package dto

type SearchStudents struct {
	FirstName   string `form:"first_name"`
	LastName    string `form:"last_name"`
	NationalID  string `form:"national_id"`
	PhoneNumber string `form:"phone_number"`
	Email       string `form:"email"`
	CampusID    string `form:"campus_id" binding:"required"`
	Page        int    `form:"page"`
	Size        int    `form:"size"`
}

type StudentResponse struct {
	ID              string `json:"id"`
	FirstNameTH     string `json:"first_name_th"`
	MiddleNameTH    string `json:"middle_name_th"`
	LastNameTH      string `json:"last_name_th"`
	FirstNameENG    string `json:"first_name_eng"`
	MiddleNameENG   string `json:"middle_name_eng"`
	LastNameENG     string `json:"last_name_eng"`
	BirthDate       string `json:"birth_date"`
	Gender          string `json:"gender"`
	NationalID      string `json:"national_id"`
	PassportID      string `json:"passport_id"`
	PhoneNumber     string `json:"phone_number"`
	Email           string `json:"email"`
	StudentType     string `json:"student_type"`
	Status          string `json:"status"`
	InactiveReason  string `json:"inactive_reson"`
	StatusUpdatedAt string `json:"Status_updated_at"`
	CreatedAt       string `json:"created_at"`
	CreatedBy       string `json:"created_by"`
	UpdatedAt       string `json:"updated_at"`
	UpdatedBy       string `json:"updated_by"`
}
