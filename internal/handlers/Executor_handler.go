package handlers

import (
	"github.com/stomatolog1/praktologia/internal/model"
	"github.com/stomatolog1/praktologia/internal/servise"
	"net/http"
	"github.com/gin-gonic/gin"
)

type ExecutorHandler struct {
	servise *servise.ExecutorServise
}

func (h *ExecutorHandler) RegisterRoutes(r *gin.Engine) {
	executor := r.Group("api/Executor")
	{
		executor.GET("", h.GetAll)
		executor.POST("", h.Create)
	}
}

func NewExecutorHandler(service *servise.ExecutorServise) *ExecutorHandler {
	return &ExecutorHandler{servise: service}
}

func (h *ExecutorHandler) GetAll(c *gin.Context) {
	executors, err := h.servise.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, executors)
}

func (h *ExecutorHandler) Create(c *gin.Context) {
	var req model.RequsetExecutor
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	executor, err := h.servise.CreateExecutor(req)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Executor name already exists"})
		return
	}
	c.JSON(http.StatusCreated, executor)
}
