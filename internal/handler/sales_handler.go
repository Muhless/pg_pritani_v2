package handler

import (
	"net/http"
	"pg_pritani/backend/internal/dto"
	"pg_pritani/backend/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type SalesHandler struct {
	service service.SalesService
}

func NewSalesHandler(service service.SalesService) *SalesHandler {
	return &SalesHandler{service}
}

func (h *SalesHandler) GetAll(ctx *gin.Context) {
	sales, err := h.service.GetAll()
	if err != nil {
		log.Error().Err(err).Msg("handler: failed to get all sales data")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to get all sales data"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": sales})
}

func (h *SalesHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Error().Err(err).Str("id", ctx.Param("id")).Msg("handler: invalid id")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	sales, err := h.service.GetByID(uint(id))
	if err != nil {
		log.Error().Err(err).Msg("handler: failed to get sales data by id")
		ctx.JSON(http.StatusNotFound, gin.H{"error": "failed to get sales data by id"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": sales})
}

func (h *SalesHandler) Create(ctx *gin.Context) {
	var req dto.CreateSalesRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("handler: failed to create sales data")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to create sales data"})
		return
	}

	if len(req.Items) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "items can't be empty"})
		return
	}

	if err := h.service.Create(req); err != nil {
		log.Error().Err(err).Msg("handler: failed to create sales data")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create sales data"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "successfully created sales data"})
}

func (h *SalesHandler) AddPayment(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Error().Err(err).Str("id", ctx.Param("id")).Msg("handler: invalid id")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req dto.UpdateSalesStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("handler: invalid json body")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body"})
		return
	}

	if err := h.service.AddPayment(uint(id), req); err != nil {
		log.Error().Err(err).Msg("handler: failed to add payment")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add payment"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully added payment"})
}

func (h *SalesHandler) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Error().Err(err).Str("id", ctx.Param("id")).Msg("handler: invalid id")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		log.Error().Err(err).Msg("handler: failed to delete sales data")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete sales data"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully deleted sales data"})
}
