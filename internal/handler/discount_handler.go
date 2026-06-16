package handler

import (
	"net/http"
	"pg_pritani/backend/internal/dto"
	"pg_pritani/backend/internal/service"
	"pg_pritani/backend/pkg/pagination"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type DiscountHandler struct {
	service service.DiscountService
}

func NewDiscountHandler(service service.DiscountService) *DiscountHandler {
	return &DiscountHandler{service}
}

func (h *DiscountHandler) GetAll(ctx *gin.Context) {
	p := pagination.GetPagination(ctx)

	discounts, total, err := h.service.GetAll(p.Page, p.Limit)
	if err != nil {
		log.Error().Err(err).Msg("handler: failed to get all discounts data")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to get all discounts data"})
		return
	}

	p.Total = float64(total)
	ctx.JSON(http.StatusOK, gin.H{"data": discounts, "pagination": p})
}

func (h *DiscountHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Error().Err(err).Str("id", ctx.Param("id")).Msg("handler: invalid id")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	discount, err := h.service.GetByID(uint(id))
	if err != nil {
		log.Error().Err(err).Msg("handler: failed to get discount data by id")
		ctx.JSON(http.StatusNotFound, gin.H{"error": "failed to get dicount data by id"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": discount})
}

func (h *DiscountHandler) Create(ctx *gin.Context) {
	var req dto.CreateDiscountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("handler: invalid json body")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body"})
		return
	}

	if err := h.service.Create(req); err != nil {
		log.Error().Err(err).Msg("handler: failed to create discount service")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create discount service"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "successfully created discount handler"})
}

func (h *DiscountHandler) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Error().Err(err).Str("id", ctx.Param("id")).Msg("handler: invalid id")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req dto.UpdateDiscountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("handler: invalid json body")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body"})
		return
	}

	if err := h.service.Update(uint(id), req); err != nil {
		log.Error().Err(err).Msg("handler: failed to update discount data")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update discount data"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully updated discount data"})
}

func (h *DiscountHandler) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Error().Err(err).Str("id", ctx.Param("id")).Msg("handler: invalid id")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		log.Error().Err(err).Msg("handler: failed to delete discount data")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete discount data"})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{"message":"successfully deleted discount data",})
}
