package dto

type Staff struct {
	ID              string
	FirstName       string
	MiddleName      string
	LastName        string
	PhoneNumber     string
	Email           string
	Status          string
	StatusUpdatedAt string
	InactiveReason  string
	RoleID          string
	CampusID        string
	CreatedAt       string
	CreatedBy       string
	UpdatedAt       string
	UpdatedBy       string
}

type CreateStaffRequest struct {
	Username      string `json:"username" binding:"required"`
	Password      string `json:"password" binding:"required"`
	FirstNameTH   string `json:"first_name_th" binding:"required"`
	MiddleNameTH  string `json:"middle_name_th"`
	LastNameTH    string `json:"last_name_th" binding:"required"`
	FirstNameENG  string `json:"first_name_eng"`
	MiddleNameENG string `json:"middle_name_eng"`
	LastNameENG   string `json:"last_name_eng"`
	PhoneNumber   string `json:"phone_number" binding:"required"`
	Email         string `json:"email" binding:"required"`
	RoleID        string `json:"role_id" binding:"required"`
	CampusID      string `json:"campus_id" binding:"required"`
}
