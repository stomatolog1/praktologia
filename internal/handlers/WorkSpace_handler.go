package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stomatolog1/praktologia/internal/model"
	"github.com/stomatolog1/praktologia/internal/servise"
)

type WorkSpaceHandler struct {
	servise *servise.WorkSpaceServise
}

func (h *WorkSpaceHandler) RegisterRoutes(r *gin.Engine) {
	workspace := r.Group("api/WorkSpace")
	{
		workspace.GET("", h.GetAll)
		workspace.POST("", h.Create)
		workspace.PUT("/:id/status", h.UpdateStatus)
		workspace.DELETE("/:id", h.Delete)
	}
}

func NewWorkSpaceHandler(service *servise.WorkSpaceServise) *WorkSpaceHandler {
	return &WorkSpaceHandler{servise: service}
}

func (h *WorkSpaceHandler) GetAll(c *gin.Context) {
	adminIDStr := c.Query("adminId")
	if adminIDStr != "" {
		adminID, err := strconv.ParseUint(adminIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid adminId"})
			return
		}
		workspaces, err := h.servise.GetByAdminID(uint(adminID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, workspaces)
		return
	}

	workspaces, err := h.servise.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, workspaces)
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

func (h *WorkSpaceHandler) UpdateStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	workspace, err := h.servise.UpdateWorkSpaceStatus(uint(id), req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, workspace)
}

func (h *WorkSpaceHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = h.servise.DeleteWorkSpace(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "workspace deleted"})
}
