package handlers

import (
	"github.com/stomatolog1/praktologia/internal/model"
	"github.com/stomatolog1/praktologia/internal/servise"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	servise *servise.CustomerServise
}

func (h *CustomerHandler) RegisterRoutes(r *gin.Engine) {
	customer := r.Group("api/Customer")
	{
		customer.GET("", h.GetAll)
		customer.POST("", h.Create)
		customer.DELETE("/:id", h.Delete)
	}
}

func NewCustomerHandler(service *servise.CustomerServise) *CustomerHandler {
	return &CustomerHandler{servise: service}
}

func (h *CustomerHandler) GetAll(c *gin.Context) {
	customers, err := h.servise.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}

func (h *CustomerHandler) Create(c *gin.Context) {
	var req model.RequestCustomer
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	customer, err := h.servise.CreateCustomer(req)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Customer name already exists"})
		return
	}
	c.JSON(http.StatusCreated, customer)
}

func (h *CustomerHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.servise.DeleteCustomer(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted"})
}
