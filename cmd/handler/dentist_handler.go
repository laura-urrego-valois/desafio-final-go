package handler

import (
	"errors"
	"net/http"
	"os"
	"proyecto_final_go/internal/domain"
	"proyecto_final_go/internal/service"
	"strings"

	"strconv"

	"github.com/gin-gonic/gin"
)

type dentistHandler struct {
	s service.DentistService
}

func NewDentistHandler(s service.DentistService) *dentistHandler {
	return &dentistHandler{
		s: s,
	}
}

// Post godoc
// @Summary Create a new dentist
// @Description This endpoint allows you to create a new dentist with the provided data.
// @Tags Dentists
// @Produce json
// @Param token header string true "TOKEN"
// @Param dentist body domain.Dentist true "Dentist"
// @Success 201 "Dentist created successfully"
// @Response 400 "Invalid dentist data or missing required fields"
// @Response 401 "Unauthorized access due to missing or invalid token"
// @Response 500 "Failed to create dentist"
// @Router /dentists [post]
func (h *dentistHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dentist domain.Dentist
		token := ctx.GetHeader("TOKEN")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token not found"})
			return
		}
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		if err := ctx.ShouldBindJSON(&dentist); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid dentist data"})
			return
		}
		if strings.TrimSpace(dentist.FirstName) == "" ||
			strings.TrimSpace(dentist.LastName) == "" ||
			strings.TrimSpace(dentist.License) == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing required fields"})
			return
		}
		err := h.s.Create(dentist)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create dentist"})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "dentist created successfully"})
	}
}

// GetByID godoc
// @Summary Get a dentist by ID
// @Description This endpoint allows you to retrieve a dentist by their ID.
// @Tags Dentists
// @Produce json
// @Param id path int true "Dentist ID"
// @Success 200 {object} domain.Dentist "Dentist"
// @Failure 400 "Invalid ID"
// @Failure 404 "Dentist not found"
// @Router /dentists/{id} [get]
func (h *dentistHandler) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New("invalid id"))
			return
		}

		dentist, err := h.s.GetByID(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "dentist not found"})
			return
		}

		ctx.JSON(http.StatusOK, dentist)
	}
}

// GetAll godoc
// @Summary Get all dentists
// @Description This endpoint allows you to retrieve all dentists.
// @Tags Dentists
// @Produce json
// @Success 200 {array} domain.Dentist "Dentists"
// @Failure 500 "Failed to retrieve dentists"
// @Router /dentists [get]
func (h *dentistHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dentists, err := h.s.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve dentists"})
			return
		}

		ctx.JSON(http.StatusOK, dentists)
	}
}

// Put godoc
// @Summary Update a dentist
// @Description This endpoint allows you to update a dentist with the provided data.
// @Tags Dentists
// @Produce json
// @Param token header string true "TOKEN"
// @Param dentist body domain.Dentist true "Updated dentist information"
// @Success 200 {object} domain.Dentist "Updated dentist"
// @Failure 400 "Invalid dentist data or missing required fields"
// @Failure 401 "Unauthorized access due to missing or invalid token"
// @Failure 500 "Failed to update dentist"
// @Router /dentists [put]
func (h *dentistHandler) Put() gin.HandlerFunc {
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

		var dentist domain.Dentist
		err := ctx.ShouldBindJSON(&dentist)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid dentist"})
			return
		}

		if dentist.Id == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "dentist id is required"})
			return
		}

		err = h.s.Update(dentist)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update dentist"})
			return
		}

		ctx.JSON(http.StatusOK, dentist)
	}
}

// Patch godoc
// @Summary Update a dentist's license
// @Description This endpoint allows you to update a dentist's license with the provided data.
// @Tags Dentists
// @Produce json
// @Param token header string true "TOKEN"
// @Param id path int true "Dentist ID"
// @Param license body string true "Updated license information"
// @Success 200 "License updated successfully"
// @Failure 400 "Invalid request or missing required fields"
// @Failure 401 "Unauthorized access due to missing or invalid token"
// @Failure 404 "Dentist not found"
// @Failure 500 "Failed to update License"
// @Router /dentists/{id} [patch]
func (h *dentistHandler) Patch() gin.HandlerFunc {
	type Request struct {
		License string `json:"license,omitempty"`
	}

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

		if r.License == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "no data provided for updating License"})
			return
		}

		oldDentist, err := h.s.GetByID(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "dentist not found"})
			return
		}

		if r.License != oldDentist.License {
			err = h.s.PatchLicense(id, r.License)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to update License"})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{"message": "License updated successfully"})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "no change detected for License"})
		}
	}
}

// Delete godoc
// @Summary Delete a dentist
// @Description This endpoint allows you to delete a dentist by their ID.
// @Tags Dentists
// @Param token header string true "TOKEN"
// @Param id path int true "Dentist ID"
// @Success 204 "Dentist deleted successfully"
// @Failure 400 "Invalid ID"
// @Failure 401 "Unauthorized access due to missing or invalid token"
// @Failure 500 "Failed to delete dentist"
// @Router /dentists/{id} [delete]
func (h *dentistHandler) Delete() gin.HandlerFunc {
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
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to delete dentist"})
			return
		}
		ctx.Status(http.StatusNoContent)
	}
}
