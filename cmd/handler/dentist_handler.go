package handler

import (
	"errors"
	"fmt"
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
// @Summary      Create a new dentist
// @Description  Create a new dentist in the system
// @Tags         dentists
// @Accept       json
// @Produce      json
// @Param        TOKEN header string true "Authorization token"
// @Param        body body domain.Dentist true "Data of the dentist to create"
// @Success      201 {object} domain.Dentist "Dentist created successfully"
// @Failure      400 {string} string "Invalid dentist data or missing required fields"
// @Failure      401 {string} string "Token not found or invalid token"
// @Failure      500 {string} string "Internal server error"
// @Router       /dentists [post]
func (h *dentistHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dentist domain.Dentist
		token := ctx.GetHeader("TOKEN")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, errors.New("invalid token"))
			return
		}
		if err := ctx.ShouldBindJSON(&dentist); err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New("invalid dentist data"))
			return
		}
		if strings.TrimSpace(dentist.FirstName) == "" ||
			strings.TrimSpace(dentist.LastName) == "" ||
			strings.TrimSpace(dentist.License) == "" {
			ctx.JSON(http.StatusBadRequest, errors.New("missing required fields"))
			return
		}
		err := h.s.Create(dentist)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			fmt.Printf(err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, dentist)
	}
}

// GetByID godoc
// @Summary      Get dentist by ID
// @Description  Get a dentist from the system by its ID
// @Tags         dentists
// @Accept       json
// @Produce      json
// @Param        id path int true "Dentist ID"
// @Success      200 {object} domain.Dentist "Dentist found successfully"
// @Failure      400 {string} string "Invalid ID"
// @Failure      404 {string} string "Dentist not found"
// @Router       /dentists/{id} [get]
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
			ctx.JSON(http.StatusNotFound, errors.New("dentist not found"))
			return
		}

		ctx.JSON(http.StatusOK, dentist)
	}
}

// GetAll godoc
// @Summary      Get all dentists
// @Description  Get all dentists from the system
// @Tags         dentists
// @Accept       json
// @Produce      json
// @Success      200 {array} domain.Dentist "List of dentists"
// @Failure      500 {string} string "Internal server error"
// @Router       /dentists [get]
func (h *dentistHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dentists, err := h.s.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errors.New("failed to retrieve dentists: "+err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, dentists)
	}
}

// Put godoc
// @Summary      Update existing dentist
// @Description  Update an existing dentist in the system
// @Tags         dentists
// @Accept       json
// @Produce      json
// @Param        TOKEN header string true "Authorization token"
// @Param        body body domain.Dentist true "Updated data of the dentist"
// @Success      200 {object} domain.Dentist "Dentist updated successfully"
// @Failure      400 {string} string "Invalid dentist data or missing dentist ID"
// @Failure      401 {string} string "Token not found or invalid token"
// @Failure      500 {string} string "Internal server error"
// @Router       /dentists [put]
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
			ctx.JSON(http.StatusBadRequest, errors.New("invalid dentist"))
			return
		}

		if dentist.Id == 0 {
			ctx.JSON(http.StatusBadRequest, errors.New("dentist ID is required"))
			return
		}

		err = h.s.Update(dentist)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errors.New(err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, dentist)
	}
}

// Patch godoc
// @Summary      Update dentist license
// @Description  Update the license of a dentist in the system
// @Tags         dentists
// @Accept       json
// @Produce      json
// @Param        TOKEN header string true "Authorization token"
// @Param        id path int true "Dentist ID"
// @Success      200 {object} domain.Dentist "Dentist updated successfully"
// @Failure      400 {string} string "Invalid ID, invalid JSON or missing license data"
// @Failure      401 {string} string "Token not found or invalid token"
// @Failure      404 {string} string "Dentist not found"
// @Failure      500 {string} string "Internal server error"
// @Router       /dentists/{id} [patch]
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
			ctx.JSON(http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		if err := ctx.ShouldBindJSON(&r); err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New("invalid json"))
			return
		}

		err = h.s.PatchLicense(id, r.License)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New(err.Error()))
			return
		}

		dentist, err := h.s.GetByID(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, errors.New("dentist not found"))
			return
		}

		ctx.JSON(http.StatusOK, dentist)
	}
}

// Delete godoc
// @Summary      Delete dentist
// @Description  Delete a dentist from the system
// @Tags         dentists
// @Accept       json
// @Produce      json
// @Param        TOKEN header string true "Authorization token"
// @Param        id path int true "Dentist ID"
// @Success      204 "Dentist deleted successfully"
// @Failure      400 {string} string "Invalid ID"
// @Failure      401 {string} string "Token not found or invalid token"
// @Failure      500 {string} string "Internal server error"
// @Router       /dentists/{id} [delete]
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
			ctx.JSON(http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New(err.Error()))
			return
		}
		ctx.Status(http.StatusNoContent)
	}
}
