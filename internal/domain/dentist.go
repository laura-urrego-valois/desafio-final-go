package domain

type Dentist struct {
	Id        int    `json:"Id"`
	FirstName string `json:"FirstName" binding:"required"`
	LastName  string `json:"LastName" binding:"required"`
	License   string `json:"License" binding:"required"`
}
