package handler

import (
	"net/http"
	"pg_pritani/backend/internal/dto"
	"pg_pritani/backend/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type EmployeeHandler struct {
	service service.EmployeeService
}

func NewEmployeeHandler(service service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service}
}

func (h *EmployeeHandler) GetAll(ctx *gin.Context) {
	employees, err := h.service.GetAll()
	if err != nil {
		log.Error().Err(err).Msg("handler: failed to get all employees data")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": employees})
}

func (h *EmployeeHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	employee, err := h.service.GetByID(uint(id))
	if err != nil {
		log.Error().Err(err).Msg("failed to get employee by id")
		ctx.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": employee})
}

func (h *EmployeeHandler) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req dto.UpdateEmployeeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "failed",
		})
		return
	}

	if err := h.service.Update(uint(id), req); err != nil {
		log.Error().Err(err).Msg("failed to update employee data")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update employee data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "employee data successfully updated"})
}

func (h *EmployeeHandler) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete employee data",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "employee data successfully deleted"})
}
