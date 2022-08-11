package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/voideus/mini-ecommerce/model"
	"github.com/voideus/mini-ecommerce/repository"
)

type ProductHandler interface {
	GetProduct(*gin.Context)
	GetAllProduct(*gin.Context)
	AddProduct(*gin.Context)
	UpdateProduct(*gin.Context)
	DeleteProduct(*gin.Context)
}

type productHandler struct {
	repo repository.ProductRepository
}

func NewProductHandler() ProductHandler {
	return &productHandler{
		repo: repository.NewProductRepository(),
	}
}

func (ph *productHandler) GetProduct(ctx *gin.Context) {
	productId := ctx.Param("id")
	prodID, err := strconv.Atoi(productId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, err := ph.repo.GetProduct(prodID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, product)
}

func (ph *productHandler) GetAllProduct(ctx *gin.Context) {
	product, err := ph.repo.GetAllProduct()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (ph *productHandler) AddProduct(ctx *gin.Context) {
	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, err := ph.repo.AddProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (ph *productHandler) UpdateProduct(ctx *gin.Context) {
	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productId := ctx.Param("id")
	prodID, err := strconv.Atoi(productId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product.ID = uint(prodID)

	product, err = ph.repo.UpdateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (ph *productHandler) DeleteProduct(ctx *gin.Context) {
	var product model.Product
	productId := ctx.Param("id")
	prodID, _ := strconv.Atoi(productId)
	product.ID = uint(prodID)

	product, err := ph.repo.DeleteProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}
