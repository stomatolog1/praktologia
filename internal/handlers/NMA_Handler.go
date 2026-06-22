package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stomatolog1/praktologia/internal/servise"
)

type NMAHandler struct {
	nmaService *servise.NMAService
}

func NewNMAHandler(nmaService *servise.NMAService) *NMAHandler {
	return &NMAHandler{nmaService: nmaService}
}

func (h *NMAHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/nma/:projectID", h.getNMA)
}

func (h *NMAHandler) getNMA(c *gin.Context) {
	projectID, err := strconv.ParseUint(c.Param("projectID"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	nma, err := h.nmaService.GetNMAData(uint(projectID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nma)
}
