package handler

import (
	"net/http"
	"pg_pritani/backend/internal/dto"
	"pg_pritani/backend/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type ProductHandler struct {
	product service.ProductService
}

func NewProductHandler(product service.ProductService) *ProductHandler {
	return &ProductHandler{product}
}

func (h *ProductHandler) GetAll(ctx *gin.Context) {
	products, err := h.product.GetAll()
	if err != nil {
		log.Error().Err(err).Msg("handler: failed to get all product")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to get all product"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": products})
}

func (h *ProductHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 24)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	product, err := h.product.GetByID(uint(id))
	if err != nil {
		log.Error().Err(err).Msg("")
		ctx.JSON(http.StatusNotFound, gin.H{"error": "failed to get product by id"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": product})
}

func (h *ProductHandler) Create(ctx *gin.Context) {
	var req dto.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to create product data"})
		return
	}

	if err := h.product.Create(req); err != nil {
		log.Error().Err(err).Msg("handler: failed to create product data")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create product data",
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "product data successfully created"})
}

func (h *ProductHandler) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 24)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req dto.UpdateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	if err := h.product.Update(uint(id), req); err != nil {
		log.Error().Err(err).Msg("handler: failed to update product data")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update product data",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "product data successfully updated"})
}

func (h *ProductHandler) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 24)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.product.Delete(uint(id)); err != nil {
		log.Error().Err(err).Msg("handler: failed to delete product data")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete product data",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "peoduct data successfully deleted"})

}
