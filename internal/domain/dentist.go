package domain

type Dentist struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	License   string `json:"license" binding:"required"`
}
