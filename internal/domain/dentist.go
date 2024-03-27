package domain

type Dentist struct {
	// @Description The unique identifier of the dentist
	// @Example 1
	Id int `json:"Id"`
	// @Description The first name of the dentist
	// @Example "Daniel"
	FirstName string `json:"FirstName" binding:"required"`
	// @Description The last name of the dentist
	// @Example "Rodríguez"
	LastName string `json:"LastName" binding:"required"`
	// @Description The license number of the dentist
	// @Example "AXMER"
	License string `json:"License" binding:"required"`
}
