package handlers

import (
	"github.com/stomatolog1/praktologia/internal/model"
	"github.com/stomatolog1/praktologia/internal/servise"
	"net/http"
	"github.com/gin-gonic/gin"

)

type AdminHandler struct {
	servise *servise.AdminServise
}

func (h *AdminHandler) RegisterRoutes(r *gin.Engine){
	Admins := r.Group("api/Admin")
	{
		Admins.POST("", h.Create)
		Admins.POST("/login", h.Login)
	}
}

func NewAdminHandler(service *servise.AdminServise) *AdminHandler{
	return &AdminHandler{servise: service}
}

func (h *AdminHandler) Login(c *gin.Context) {
	var req model.RequestAdminLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	admin, err := h.servise.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, admin)
}

func (h *AdminHandler) Create(c *gin.Context) {
	var req model.RequestAdminAkk
	if err:= c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	admin, err := h.servise.CreateAdmin(req)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, model.AdminResponse{ID: admin.ID, Login: admin.Login})
}