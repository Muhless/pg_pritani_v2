package handler

import (
	"net/http"
	"pg_pritani/backend/internal/dto"
	"pg_pritani/backend/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type PaymentHandler struct {
	service service.PaymentService
}

func NewPaymentHandler(service service.PaymentService) *PaymentHandler {
	return &PaymentHandler{service}
}

func (h *PaymentHandler) GetAll(ctx *gin.Context) {
	payments, err := h.service.GetAll()
	if err != nil {
		log.Error().Err(err).Msg("handler: failed to get all payments data")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get all payments data"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": payments})
}

func (h *PaymentHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Error().Err(err).Str("id", ctx.Param("id")).Msg("handler: invalid id")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	payment, err := h.service.GetByID(uint(id))
	if err != nil {
		log.Error().Err(err).Uint64("id", id).Msg("handler: failed to get payment data by id")
		ctx.JSON(http.StatusNotFound, gin.H{"error": "failed to get payment data by id"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": payment})
}

func (h *PaymentHandler) GetBySalesID(ctx *gin.Context) {
	salesID, err := strconv.ParseUint(ctx.Param("sales_id"), 10, 64)
	if err != nil {
		log.Error().Err(err).Str("sales_id", ctx.Param("sales_id")).Msg("handler: invalid sales_id")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid sales_id"})
		return
	}

	payments, err := h.service.GetBySalesID(uint(salesID))
	if err != nil {
		log.Error().Err(err).Uint64("sales_id", salesID).Msg("failed to get payments data by sales_id")
		ctx.JSON(http.StatusNotFound, gin.H{"error": "failed to get payments data by sales_id"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": payments})
}

func (h *PaymentHandler) Create(ctx *gin.Context) {
	var req dto.CreatePaymentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("handler: invalid json body")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body"})
		return
	}

	if req.Amount <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "amount must be greater than 0"})
		return
	}

	if err := h.service.Create(req); err != nil {
		log.Error().Err(err).Msg("handler: failed to create payment data")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create payment data"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "successfully created payment data"})
}

func (h *PaymentHandler) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Error().Err(err).Str("id", ctx.Param("id")).Msg("handler: invalid id")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		log.Error().Err(err).Uint64("id", id).Msg("handler: failed to delete payment data by id")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete payment data by id"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully deleted payment data"})
}
