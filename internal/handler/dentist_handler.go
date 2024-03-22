package handler

import (
	"net/http"
	"os"
	"proyecto_final_go/internal/domain"
	"proyecto_final_go/internal/service"

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

func (h *dentistHandler) Post() gin.HandlerFunc {
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

		var dentist domain.Dentist
		if err := ctx.ShouldBindJSON(&dentist); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid dentist data"})
			return
		}

		if dentist.FirstName == "" || dentist.LastName == "" || dentist.License == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing required fields"})
			return
		}

		createdDentist, err := h.s.Create(&dentist)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create dentist"})
			return
		}

		ctx.JSON(http.StatusCreated, createdDentist)
	}
}

func (h *dentistHandler) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
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

func (h *dentistHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		dentists, err := h.s.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get dentists"})
			return
		}

		ctx.JSON(http.StatusOK, dentists)
	}
}

func (h *dentistHandler) Put() gin.HandlerFunc {
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

		var dentist domain.Dentist
		err := ctx.ShouldBindJSON(&dentist)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid dentist"})
			return
		}

		err = h.s.Update(&dentist)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, dentist)
	}
}

func (h *dentistHandler) Patch() gin.HandlerFunc {
	type Request struct {
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
		License   string `json:"license,omitempty"`
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

		update := domain.Dentist{
			FirstName: r.FirstName,
			LastName:  r.LastName,
			License:   r.License,
		}

		dentist, err := h.s.GetByID(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "dentist not found"})
			return
		}

		if update.FirstName == "" && update.LastName == "" && update.License == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "no fields to update"})
			return
		}

		if update.FirstName != "" {
			dentist.FirstName = update.FirstName
		}
		if update.LastName != "" {
			dentist.LastName = update.LastName
		}
		if update.License != "" {
			dentist.License = update.License
		}

		if err := h.s.Update(dentist); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, dentist)
	}
}

func (h *dentistHandler) Delete() gin.HandlerFunc {
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
		ctx.JSON(http.StatusNoContent, gin.H{"msg": "dentist deleted"})
	}
}
