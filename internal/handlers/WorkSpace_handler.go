package handlers

import (
	"github.com/stomatolog1/praktologia/internal/model"
	"github.com/stomatolog1/praktologia/internal/servise"
	"net/http"
	"github.com/gin-gonic/gin"
)

type WorkSpaceHandler struct {
	servise *servise.WorkSpaceServise
}

func (h *WorkSpaceHandler) RegisterRoutes(r *gin.Engine) {
	workspace := r.Group("api/WorkSpace")
	{
		workspace.POST("", h.Create)
	}
}

func NewWorkSpaceHandler(service *servise.WorkSpaceServise) *WorkSpaceHandler {
	return &WorkSpaceHandler{servise: service}
}

func (h *WorkSpaceHandler) Create(c *gin.Context) {
	var req model.RequestWorkSpace
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	workspace, err := h.servise.CreateWorkSpace(req)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "WorkSpace name already exists"})
		return
	}
	c.JSON(http.StatusCreated, workspace)
}
