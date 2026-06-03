package handler

import (
	"net/http"
	"pg_pritani/backend/internal/dto"
	"pg_pritani/backend/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type SupplierHandler struct {
	service service.SupplierService
}

func NewSupplierHandler(service service.SupplierService) *SupplierHandler {
	return &SupplierHandler{service}
}

func (h *SupplierHandler) GetAll(ctx *gin.Context) {
	suppliers, err := h.service.GetAll()
	if err != nil {
		log.Error().Err(err).Msg("handler: failed to get suppliers data")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to get suppliers data"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": suppliers})
}

func (h *SupplierHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 24)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	supplier, err := h.service.GetByID(uint(id))
	if err != nil {
		log.Error().Err(err).Msg("failed to get supplier data by id")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get supplier data by id"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": supplier})
}

func (h *SupplierHandler) Create(ctx *gin.Context) {
	var req dto.CreateSupplierRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("handler: invalid request body")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	if err := h.service.Create(req); err != nil {
		log.Error().Err(err).Msg("handler: failed to create supplier data")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create supplier data",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "supplier data successfully created"})
}

func (h *SupplierHandler) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 24)
	if err != nil {
		log.Error().Err(err).Msg("invalid id")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
	}

	var req dto.UpdateSupplierRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("handler: invalid request body")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	if err := h.service.Update(uint(id), req); err != nil {
		log.Error().Err(err).Msg("handler: failed to update supplier data")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update supplier data",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "supplier data successfully updated"})
}

func (h *SupplierHandler) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 24)
	if err != nil {
		log.Error().Err(err).Msg("handler: invalid id")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
	}

	if err := h.service.Delete(uint(id)); err != nil {
		log.Error().Err(err).Msg("handler: failed to delete supplier data")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete supplier data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error": "supplier data successfully deleted"})
}
