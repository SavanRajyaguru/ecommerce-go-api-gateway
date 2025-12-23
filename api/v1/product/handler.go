package product

import (
	"ecommerce-go-api-gateway/models"
	"ecommerce-go-api-gateway/pkg/utils"
	"ecommerce-go-api-gateway/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service services.ProductService
}

func NewProductHandler(service services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	products, err := h.service.ListProducts()
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to list products", err.Error())
		return
	}

	utils.SendSuccess(c, http.StatusOK, "Products list", products)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid product ID", err.Error())
		return
	}

	product, err := h.service.GetProduct(uint(id))
	if err != nil {
		utils.SendError(c, http.StatusNotFound, "Product not found", err.Error())
		return
	}

	utils.SendSuccess(c, http.StatusOK, "Product details", product)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req models.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	product, err := h.service.CreateProduct(req)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to create product", err.Error())
		return
	}

	utils.SendSuccess(c, http.StatusCreated, "Product created successfully", product)
}
