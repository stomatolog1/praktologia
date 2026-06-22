package handlers

import (
	"github.com/stomatolog1/praktologia/internal/model"
	"github.com/stomatolog1/praktologia/internal/servise"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type SotrudnikHandler struct {
	servise *servise.SotrudnikServise
}

func (h *SotrudnikHandler) RegisterRoutes(r *gin.Engine) {
	sotrudnik := r.Group("api/Sotrudnik")
	{
		sotrudnik.GET("", h.GetAll)
		sotrudnik.POST("", h.Create)
		sotrudnik.DELETE("/:id", h.Delete)
	}
}

func NewSotrudnikHandler(service *servise.SotrudnikServise) *SotrudnikHandler {
	return &SotrudnikHandler{servise: service}
}

func (h *SotrudnikHandler) GetAll(c *gin.Context) {
	items, err := h.servise.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *SotrudnikHandler) Create(c *gin.Context) {
	var req model.RequestSotrudnik
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	sotrudnik, err := h.servise.CreateSotrudnik(req)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Sotrudnik name already exists"})
		return
	}
	c.JSON(http.StatusCreated, sotrudnik)
}

func (h *SotrudnikHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.servise.DeleteSotrudnik(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sotrudnik not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sotrudnik deleted"})
}
