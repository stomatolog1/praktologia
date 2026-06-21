package handlers

import (
	"github.com/stomatolog1/praktologia/internal/model"
	"github.com/stomatolog1/praktologia/internal/servise"
	"net/http"
	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	servise *servise.ProjectServise
}

func (h *ProjectHandler) RegisterRoutes(r *gin.Engine) {
	project := r.Group("api/Project")
	{
		project.POST("", h.Create)
	}
}

func NewProjectHandler(service *servise.ProjectServise) *ProjectHandler {
	return &ProjectHandler{servise: service}
}

func (h *ProjectHandler) Create(c *gin.Context) {
	var req model.RequestProject
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	project, err := h.servise.CreateProject(req)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Project name already exists"})
		return
	}
	c.JSON(http.StatusCreated, project)
}
