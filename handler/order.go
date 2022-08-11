package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/voideus/mini-ecommerce/repository"
)

type OrderHandler interface {
	OrderProduct(*gin.Context)
}

type orderHandler struct {
	repo repository.OrderRepository
}

func NewOrderHandler() OrderHandler {
	return &orderHandler{
		repo: repository.NewOrderRepository(),
	}
}

func (oh *orderHandler) OrderProduct(ctx *gin.Context) {
	prodIDStr := ctx.Param("product")
	if prodID, err := strconv.Atoi(prodIDStr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		quantityIDStr := ctx.Param("quantity")
		if quantityID, err := strconv.Atoi(quantityIDStr); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			userID := ctx.GetFloat64("userID")
			if err := oh.repo.OrderProduct(int(userID), prodID, quantityID); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.String(http.StatusOK, "Product Successfully Ordered.")
			}
		}
	}

}
