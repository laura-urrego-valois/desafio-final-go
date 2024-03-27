package handler

import (
	"errors"
	"net/http"
	"os"
	"proyecto_final_go/internal/domain"
	"proyecto_final_go/internal/service"
	"strconv"
	"strings"

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

// Post godoc
// @Summary Create a new patient
// @Description This endpoint allows you to create a new patient with the provided data.
// @Tags Patients
// @Produce json
// @Param token header string true "TOKEN"
// @Param patient body domain.Patient true "Patient"
// @Success 201 "Patient created successfully"
// @Response 400 "Invalid patient data or missing required fields"
// @Response 401 "Unauthorized access due to missing or invalid token"
// @Response 500 "Failed to create patient"
// @Router /patients [post]
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

		if strings.TrimSpace(patient.FirstName) == "" ||
			strings.TrimSpace(patient.LastName) == "" ||
			strings.TrimSpace(patient.Address) == "" ||
			strings.TrimSpace(patient.DNI) == "" ||
			strings.TrimSpace(patient.ReleaseDate) == "" {
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

// GetByID godoc
// @Summary Get a patient by ID
// @Description This endpoint allows you to retrieve a patient by their ID.
// @Tags Patients
// @Produce json
// @Param id path int true "Patient ID"
// @Success 200 {object} domain.Patient "Patient"
// @Failure 400 "Invalid ID"
// @Failure 404 "Patient not found"
// @Router /patients/{id} [get]
func (h *patientHandler) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New("invalid id"))
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

// GetAll godoc
// @Summary Get all patients
// @Description This endpoint allows you to retrieve all patients.
// @Tags Patients
// @Produce json
// @Success 200 {array} domain.Patient "Patients"
// @Failure 500 "Failed to retrieve patients"
// @Router /dentists [get]
func (h *patientHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		patients, err := h.s.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve patients"})
			return
		}

		ctx.JSON(http.StatusOK, patients)
	}
}

// Put godoc
// @Summary Update a patient
// @Description This endpoint allows you to update a patient with the provided data.
// @Tags Patients
// @Produce json
// @Param token header string true "TOKEN"
// @Param patient body domain.Patient true "Updated patient information"
// @Success 200 {object} domain.Patient "Updated patient"
// @Failure 400 "Invalid patient data or missing required fields"
// @Failure 401 "Unauthorized access due to missing or invalid token"
// @Failure 500 "Failed to update patient"
// @Router /patients [put]
func (h *patientHandler) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("TOKEN")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, errors.New("invalid token"))
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
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to update patient"})
			return
		}

		ctx.JSON(http.StatusOK, patient)
	}
}

// Patch godoc
// @Summary Update a patient's address
// @Description This endpoint allows you to update a patient's address with the provided data.
// @Tags Patients
// @Produce json
// @Param token header string true "TOKEN"
// @Param id path int true "Patient ID"
// @Param address body string true "Updated address information"
// @Success 200 "Address updated successfully"
// @Failure 400 "Invalid request or missing required fields"
// @Failure 401 "Unauthorized access due to missing or invalid token"
// @Failure 404 "Patient not found"
// @Failure 500 "Failed to update Address"
// @Router /patients/{id} [patch]
func (h *patientHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Address string `json:"address,omitempty"`
	}

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

		var r Request
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		if err := ctx.ShouldBindJSON(&r); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
			return
		}

		if r.Address == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "no data provided for updating Address"})
			return
		}

		oldPatient, err := h.s.GetByID(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "patient not found"})
			return
		}

		if r.Address != oldPatient.Address {
			err = h.s.PatchAddress(id, r.Address)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to update Address"})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "Address updated successfully"})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "no change detected for Address"})
		}
	}
}

// Delete godoc
// @Summary Delete a patient
// @Description This endpoint allows you to delete a patient by their ID.
// @Tags Patients
// @Param token header string true "TOKEN"
// @Param id path int true "Patient ID"
// @Success 204 "Patient deleted successfully"
// @Failure 400 "Invalid ID"
// @Failure 401 "Unauthorized access due to missing or invalid token"
// @Failure 500 "Failed to delete patient"
// @Router /patients/{id} [delete]
func (h *patientHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("TOKEN")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, errors.New("invalid token"))
			return
		}

		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New("failed to delete patient"))
			return
		}
		ctx.Status(http.StatusNoContent)
	}
}
