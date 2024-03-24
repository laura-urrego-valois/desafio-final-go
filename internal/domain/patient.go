package domain

type Patient struct {
	Id          int    `json:"Id"`
	FirstName   string `json:"FirstName" binding:"required"`
	LastName    string `json:"LastName" binding:"required"`
	Address     string `json:"Address" binding:"required"`
	DNI         string `json:"DNI" binding:"required"`
	ReleaseDate string `json:"ReleaseDate" binding:"required"`
}
