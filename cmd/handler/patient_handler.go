package handler

import (
	"net/http"
	"os"
	"proyecto_final_go/internal/domain"
	"proyecto_final_go/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// -------------------------------------
type patientHandler struct {
	s service.PatientService
}

func NewPatientHandler(s service.PatientService) *patientHandler {
	return &patientHandler{
		s: s,
	}
}

//--------------------------------------

func (h *patientHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("TOKEN")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token not found"})
			return
		}
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		var patient domain.Patient
		if err := ctx.ShouldBindJSON(&patient); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient data"})
			return
		}

		if patient.FirstName == "" || patient.LastName == "" || patient.Address == "" || patient.DNI == "" || patient.ReleaseDate == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing required fields"})
			return
		}

		err := h.s.Create(patient)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create patient"})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "Patient created successfully"})
	}
}

func (h *patientHandler) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("TOKEN")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token not found"})
			return
		}
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		idParam := ctx.Param("Id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		patient, err := h.s.GetByID(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "patient not found"})
			return
		}

		ctx.JSON(http.StatusOK, patient)
	}
}

func (h *patientHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		patients, err := h.s.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get patients"})
			return
		}

		ctx.JSON(http.StatusOK, patients)
	}
}

func (h *patientHandler) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("TOKEN")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token not found"})
			return
		}
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		var patient domain.Patient
		err := ctx.ShouldBindJSON(&patient)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient"})
			return
		}

		err = h.s.Update(patient)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, patient)
	}
}

func (h *patientHandler) Patch() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var json struct {
			Address string `json:"address" binding:"required"`
		}

		if err := c.BindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = h.s.PatchAddress(id, json.Address)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Address patched successfully"})
	}
}

func (h *patientHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("TOKEN")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token not found"})
			return
		}
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusNoContent, gin.H{"msg": "patient deleted"})
	}
}
