package handler

import (
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
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if appointment.Patient == (domain.Patient{}) && appointment.Dentist == (domain.Dentist{}) && appointment.Date == "" && appointment.Hour == "" && appointment.Description == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing required fields"})
		}
		err := h.appointmentService.Create(appointment)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, appointment)
	}
}

func (h *appointmentHandler) PostByDNIAndLicese() gin.HandlerFunc {
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

func (h *appointmentHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
			return
		}
		appointment, err := h.appointmentService.GetByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, appointment)
	}
}

func (h *appointmentHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		appointments, err := h.appointmentService.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, appointments)
	}
}

func (h *appointmentHandler) Put() gin.HandlerFunc {
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
		err := c.ShouldBindJSON(&appointment)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid appointment id"})
			return
		}
		if appointment.Id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "appointment id is required"})
			return
		}
		if err := c.ShouldBindJSON(&appointment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = h.appointmentService.Update(appointment)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, appointment)
	}
}

func (h *appointmentHandler) Delete() gin.HandlerFunc {
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
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment Id"})
			return
		}
		err = h.appointmentService.Delete(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusNoContent, gin.H{"msg": "appointment deleted successfully"})
	}
}

func (h *appointmentHandler) GetByPatientDNI() gin.HandlerFunc {
	return func(c *gin.Context) {
		dni := c.Query("dni")
		appointments, err := h.appointmentService.GetByPatientDNI(dni)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, appointments)
	}
}

func (h *appointmentHandler) PatchDescription() gin.HandlerFunc {
	type Request struct {
		Description string `json:"description,omitempty"`
	}
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
		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment id"})
			return
		}
		if err := c.BindJSON(&r); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = h.appointmentService.PatchDescription(id, r.Description)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
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
