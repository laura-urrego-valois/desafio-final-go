package handler

import (
	"net/http"
	"strconv"

	"proyecto_final_go/internal/domain"
	"proyecto_final_go/internal/service"

	"github.com/gin-gonic/gin"
)

type AppointmentHandler struct {
	appointmentService service.AppointmentService
}

func NewAppointmentHandler(appointmentService service.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{appointmentService}
}

func (h *AppointmentHandler) Create(c *gin.Context) {
	var appointment domain.Appointment
	if err := c.BindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.appointmentService.Create(appointment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Appointment created successfully"})
}

func (h *AppointmentHandler) CreateByPatientDNIAndDentistLicense(c *gin.Context) {
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

func (h *AppointmentHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
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

func (h *AppointmentHandler) GetAll(c *gin.Context) {
	appointments, err := h.appointmentService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointments)
}

func (h *AppointmentHandler) Update(c *gin.Context) {
	var appointment domain.Appointment
	if err := c.BindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}
	appointment.Id = id

	if err := h.appointmentService.Update(appointment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Appointment updated successfully"})
}

func (h *AppointmentHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}

	if err := h.appointmentService.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Appointment deleted successfully"})
}

func (h *AppointmentHandler) GetByPatientDNI(c *gin.Context) {
	dni := c.Query("dni")
	appointments, err := h.appointmentService.GetByPatientDNI(dni)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, appointments)
}

func (h *AppointmentHandler) PatchDescription(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}

	var req struct {
		Description string `json:"description" binding:"required"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.appointmentService.PatchDescription(id, req.Description); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Appointment description updated successfully"})
}
