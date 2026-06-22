package handlers

import (
	"github.com/stomatolog1/praktologia/internal/model"
	"github.com/stomatolog1/praktologia/internal/servise"
	"net/http"
	"github.com/gin-gonic/gin"
)

type EquipmentHandler struct {
	servise *servise.EquipmentServise
}

func (h *EquipmentHandler) RegisterRoutes(r *gin.Engine) {
	equipment := r.Group("api/Equipment")
	{
		equipment.GET("", h.GetAll)
		equipment.POST("", h.Create)
	}
}

func NewEquipmentHandler(service *servise.EquipmentServise) *EquipmentHandler {
	return &EquipmentHandler{servise: service}
}

func (h *EquipmentHandler) GetAll(c *gin.Context) {
	items, err := h.servise.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *EquipmentHandler) Create(c *gin.Context) {
	var req model.RequsetEquipment
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	equipment, err := h.servise.CreateEquipment(req)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Equipment name already exists"})
		return
	}
	c.JSON(http.StatusCreated, equipment)
}
