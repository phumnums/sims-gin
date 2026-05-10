package dto

type University struct {
	ID        string
	Name      string
	CreatedAt string
	CreatedBy string
	UpdatedAt string
	UpdatedBy string
}

type campus struct {
	ID              string
	Name            string
	Status          string
	StatusUpdatedAt string
	UniversityID    string
	CreatedAt       string
	CreatedBy       string
	UpdatedAt       string
	UpdatedBy       string
}
