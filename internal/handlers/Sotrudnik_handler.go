package handlers

import (
	"github.com/stomatolog1/praktologia/internal/model"
	"github.com/stomatolog1/praktologia/internal/servise"
	"net/http"
	"github.com/gin-gonic/gin"
)

type SotrudnikHandler struct {
	servise *servise.SotrudnikServise
}

func (h *SotrudnikHandler) RegisterRoutes(r *gin.Engine) {
	sotrudnik := r.Group("api/Sotrudnik")
	{
		sotrudnik.POST("", h.Create)
	}
}

func NewSotrudnikHandler(service *servise.SotrudnikServise) *SotrudnikHandler {
	return &SotrudnikHandler{servise: service}
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
