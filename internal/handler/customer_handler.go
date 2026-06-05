package handler

import (
	"net/http"
	"pg_pritani/backend/internal/dto"
	"pg_pritani/backend/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type CustomerHandler struct {
	service service.CustomerService
}

func NewCustomerHandler(service service.CustomerService) *CustomerHandler {
	return &CustomerHandler{service}
}

func (h *CustomerHandler) GetAll(ctx *gin.Context) {
	customers, err := h.service.GetAll()
	if err != nil {
		log.Error().Err(err).Msg("handler: failed to get all customer data")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to get all customer data"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": customers})
}

func (h *CustomerHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 24)
	if err != nil {
		log.Error().Err(err).Str("id", ctx.Param("id")).Msg("handler: invalid id")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	customer, err := h.service.GetByID(uint(id))
	if err != nil {
		log.Error().Err(err).Msg("handler: failed to get customer data by id")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get customer data by id"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": customer})
}

func (h *CustomerHandler) Create(ctx *gin.Context) {
	var req dto.CreateCustomerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("handler: invalid requset body")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := h.service.Create(req); err != nil {
		log.Error().Err(err).Msg("handler: failed to create customer data")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create customer data"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "successfully created customer data"})
}

func (h *CustomerHandler) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Error().Err(err).Str("id", ctx.Param("id")).Msg("handler: invalid id")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req dto.UpdateCustomerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("handler: invalid json body")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body"})
		return
	}

	if err := h.service.Update(uint(id), req); err != nil {
		log.Error().Err(err).Msg("handler: failed to update customer data")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update customer data"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully updated customer data"})
}

func (h *CustomerHandler) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Error().Err(err).Str("id", ctx.Param("id")).Msg("handler: invalid id")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		log.Error().Err(err).Msg("handler: failed to delete customer data")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete customer data"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully deleted customer data"})
}
