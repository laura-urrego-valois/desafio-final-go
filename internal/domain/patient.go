package domain

type Patient struct {
	// @Description The unique identifier of the patient
	// @Example 1
	Id int `json:"Id"`
	// @Description The first name of the patient
	// @Example "Daniel"
	FirstName string `json:"FirstName" binding:"required"`
	// @Description The last name of the patient
	// @Example "Rodríguez"
	LastName string `json:"LastName" binding:"required"`
	// @Description The address of the patient
	// @Example "Av. 22 # 40"
	Address string `json:"Address" binding:"required"`
	// @Description The DNI of the patient
	// @Example "538434"
	DNI string `json:"DNI" binding:"required"`
	// @Description The release date of the patient (dd/MM/YYYY)
	// @Example "30/03/2024"
	ReleaseDate string `json:"ReleaseDate" binding:"required"`
}
