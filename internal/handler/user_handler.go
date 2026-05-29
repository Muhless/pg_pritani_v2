package handler

import (
	"net/http"
	"pg_pritani/backend/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) GetAll(ctx *gin.Context) {
	users, err := h.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

func (h *UserHandler) GetByID (ctx *gin.Context)  {
	id,err:= strconv.Atoi()
user,err := h.service.GetByID()
if err != nil {
	return nil, err
}
}
