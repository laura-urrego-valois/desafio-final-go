package domain

type Appointment struct {
	ID          int     `json:"id"`
	Patient     Patient `json:"patient" binding:"required"`
	Dentist     Dentist `json:"dentist" binding:"required"`
	Date        string  `json:"date" binding:"required"`
	Hour        string  `json:"hour" binding:"required"`
	Description string  `json:"description" binding:"required"`
}
