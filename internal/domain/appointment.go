package domain

type Appointment struct {
	// @Description The unique identifier of the appointment
	// @Example 1
	Id int `json:"Id"`
	// @Description Information related to the patient
	Patient Patient `json:"patients_Id" binding:"required"`
	// @Description Information related to the patient
	Dentist Dentist `json:"dentists_Id" binding:"required"`
	// @Description The release date of the patient (dd/MM/YYYY)
	// @Example "30/03/2024"
	Date string `json:"Date" binding:"required"`
	// @Description The time of the appointment in 24-hour format
	// @Example "09:00"
	Hour string `json:"Hour" binding:"required"`
	// @Description The description of the appointment
	// @Example "Routine checkup"
	Description string `json:"Description" binding:"required"`
}
