package domain

type Patient struct {
	Id          int    `json:"id"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	DNI         string `json:"dni" binding:"required"`
	ReleaseDate string `json:"release_date" binding:"required"`
}
