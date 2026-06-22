package handlers

import (
	"github.com/stomatolog1/praktologia/internal/model"
	"github.com/stomatolog1/praktologia/internal/servise"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	servise *servise.ProjectServise
}

func (h *ProjectHandler) RegisterRoutes(r *gin.Engine) {
	project := r.Group("api/Project")
	{
		project.GET("", h.GetAll)
		project.POST("", h.Create)
	}
}

func NewProjectHandler(service *servise.ProjectServise) *ProjectHandler {
	return &ProjectHandler{servise: service}
}

func (h *ProjectHandler) GetAll(c *gin.Context) {
	workspaceIDStr := c.Query("WorkSpaceID")
	if workspaceIDStr != "" {
		workspaceID, err := strconv.ParseUint(workspaceIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid workspaceId"})
			return
		}
		projects, err := h.servise.GetByWorkSpace(uint(workspaceID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, projects)
		return
	}

	projects, err := h.servise.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, projects)
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
