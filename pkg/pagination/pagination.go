package pagination

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type pagination struct {
	Page  int     `json:"page"`
	Limit int     `json:"limit"`
	Total float64 `json:"total"`
}

func GetPagination(ctx *gin.Context) pagination {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	return pagination{
		Page:  page,
		Limit: limit,
	}
}

func (p *pagination) Offset() int {
	return (p.Page - 1) * p.Limit
}
