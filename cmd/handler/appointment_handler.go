package handler

import (
	"errors"
	"net/http"
	"os"
	"strconv"

	"proyecto_final_go/internal/domain"
	"proyecto_final_go/internal/service"

	"github.com/gin-gonic/gin"
)

type appointmentHandler struct {
	appointmentService service.AppointmentService
}

func NewAppointmentHandler(appointmentService service.AppointmentService) *appointmentHandler {
	return &appointmentHandler{appointmentService}
}

// Post godoc
// @Summary Create a new appointment
// @Description This endpoint allows you to create a new appointment with the provided data.
// @Tags Appointments
// @Produce json
// @Param token header string true "TOKEN"
// @Param dentist body domain.Appointment true "Appointment"
// @Success 201 "Appointment created successfully"
// @Response 400 "Invalid appointment data or missing required fields"
// @Response 401 "Unauthorized access due to missing or invalid token"
// @Response 500 "Failed to create appointment"
// @Router /appointments [post]
func (h *appointmentHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token not found"})
			return
		}
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		var appointment domain.Appointment
		if err := c.ShouldBindJSON(&appointment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error processing appointment data. Please ensure you send valid JSON with all required fields."})
			return
		}
		if appointment.Patient == (domain.Patient{}) ||
			appointment.Dentist == (domain.Dentist{}) ||
			appointment.Date == "" ||
			appointment.Hour == "" ||
			appointment.Description == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing required fields"})
			return
		}
		err := h.appointmentService.Create(appointment)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating appointment"})
			return
		}

		c.Status(http.StatusCreated)
	}
}

// PostByDNIAndLicese godoc
// @Summary      Create a new appointment by patient DNI and dentist license
// @Description  Create a new appointment in the system by patient DNI and dentist license
// @Tags         Appointments
// @Produce      json
// @Param token header string true "TOKEN"
// @Param        patient_dni query string true "Patient DNI"
// @Param        license query string true "Dentist license"
// @Param        date query string true "Appointment date"
// @Param        hour query string true "Appointment hour"
// @Param        description query string true "Appointment description"
// @Success      201 {object} domain.Appointment "Appointment created successfully"
// @Failure      400 "Invalid parameters or missing required fields"
// @Failure      401 "Token not found or invalid token"
// @Failure      500 "Internal server error"
// @Router       /appointments/dnilicense [post]
func (h *appointmentHandler) PostByDNIAndLicense() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token not found"})
			return
		}
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		patientDNI := c.Query("patient_dni")
		license := c.Query("license")
		date := c.Query("date")
		hour := c.Query("hour")
		description := c.Query("description")

		appointments, err := h.appointmentService.CreateByPatientDNIAndDentistLicense(patientDNI, license, date, hour, description)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, appointments)
	}
}

// GetByID godoc
// @Summary Get an appointment by ID
// @Description This endpoint allows you to retrieve an appointment by its ID.
// @Tags Appointments
// @Produce json
// @Param id path int true "Appointment ID"
// @Success 200 {object} domain.Appointment "Appointment"
// @Failure 400 "Invalid ID"
// @Failure 404 "Appointment not found"
// @Router /appointments/{id} [get]
func (h *appointmentHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, errors.New("invalid appointment id"))
			return
		}
		appointment, err := h.appointmentService.GetByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "appointment not found"})
			return
		}

		c.JSON(http.StatusOK, appointment)
	}
}

// GetAll godoc
// @Summary Get all appointments
// @Description This endpoint allows you to retrieve all appointments.
// @Tags Appointments
// @Produce json
// @Success 200 {array} domain.Appointment "Appointments"
// @Failure 500 "Failed to retrieve appointments"
// @Router /appointments [get]
func (h *appointmentHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		appointments, err := h.appointmentService.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve all appointments."})
			return
		}

		c.JSON(http.StatusOK, appointments)
	}
}

// Put godoc
// @Summary Update an appointment
// @Description This endpoint allows you to update an appointment with the provided data.
// @Tags Appointments
// @Produce json
// @Param token header string true "TOKEN"
// @Param dentist body domain.Appointment true "Updated appointment information"
// @Success 200 {object} domain.Appointment "Updated appointment"
// @Failure 400 "Invalid appointment data or missing required fields"
// @Failure 401 "Unauthorized access due to missing or invalid token"
// @Failure 500 "Failed to update appointment"
// @Router /appointments [put]
func (h *appointmentHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			c.JSON(http.StatusUnauthorized, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, errors.New("invalid token"))
			return
		}

		var appointment domain.Appointment
		if err := c.ShouldBindJSON(&appointment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid appointment data"})
			return
		}

		if appointment.Id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "appointment id is required"})
			return
		}

		err := h.appointmentService.Update(appointment)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update appointment"})
			return
		}

		c.JSON(http.StatusOK, appointment)
	}
}

// Delete godoc
// @Summary Delete an appointment
// @Description This endpoint allows you to delete an appointment by its ID.
// @Tags Appointments
// @Param token header string true "TOKEN"
// @Param id path int true "Appointment ID"
// @Success 204 "Appointment deleted successfully"
// @Failure 400 "Invalid ID"
// @Failure 401 "Unauthorized access due to missing or invalid token"
// @Failure 500 "Failed to delete appointment"
// @Router /appointments/{id} [delete]
func (h *appointmentHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			c.JSON(http.StatusUnauthorized, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, errors.New("invalid token"))
			return
		}
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, errors.New("invalid appointment id"))
			return
		}
		err = h.appointmentService.Delete(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, errors.New("failed to delete appointment"))
			return
		}

		c.Status(http.StatusNoContent)
	}
}

// GetByPatientDNI godoc
// @Summary Get appointments by patient DNI
// @Description This endpoint allows you to retrieve appointments for a patient by their DNI.
// @Tags Appointments
// @Param dni query string true "Patient DNI"
// @Produce json
// @Success 200 {array} domain.Appointment "Appointments"
// @Failure 400 "DNI parameter is required"
// @Failure 404 "No appointments found for this patient DNI"
// @Router /appointments/patient [get]
func (h *appointmentHandler) GetByPatientDNI() gin.HandlerFunc {
	return func(c *gin.Context) {
		dni := c.Query("dni")
		if dni == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "DNI parameter is required"})
			return
		}
		appointments, err := h.appointmentService.GetByPatientDNI(dni)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "no appointments found for this patient DNI"})
			return
		}
		c.JSON(http.StatusOK, appointments)
	}
}

// Patch godoc
// @Summary Update an appointment's description
// @Description This endpoint allows you to update a appointment's description with the provided data.
// @Tags Appointments
// @Produce json
// @Param token header string true "TOKEN"
// @Param id path int true "Appointment ID"
// @Param description body string true "Updated description information"
// @Success 200 "Description updated successfully"
// @Failure 400 "Invalid request or missing required fields"
// @Failure 401 "Unauthorized access due to missing or invalid token"
// @Failure 404 "Appointment not found"
// @Failure 500 "Failed to update Description"
// @Router /appointments/{id} [patch]
func (h *appointmentHandler) PatchDescription() gin.HandlerFunc {
	type Request struct {
		Description string `json:"description,omitempty"`
	}

	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			c.JSON(http.StatusUnauthorized, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, errors.New("invalid token"))
			return
		}

		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, errors.New("invalid appointment id"))
			return
		}
		if err := c.BindJSON(&r); err != nil {
			c.JSON(http.StatusBadRequest, errors.New("invalid JSON data"))
			return
		}

		if r.Description == "" {
			c.JSON(http.StatusBadRequest, errors.New("no data provided for updating description"))
			return
		}

		oldAppointment, err := h.appointmentService.GetByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, errors.New("appointment not found"))
			return
		}

		if r.Description != oldAppointment.Description {
			err = h.appointmentService.PatchDescription(id, r.Description)
			if err != nil {
				c.JSON(http.StatusBadRequest, errors.New("failed to update description"))
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "Description updated successfully"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "no change detected for description"})
		}
	}
}
