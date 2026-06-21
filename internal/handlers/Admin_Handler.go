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
	}
}

func NewAdminHandler(service *servise.AdminServise) *AdminHandler{
	return &AdminHandler{servise: service}
}

func (h *AdminHandler) Create(c *gin.Context){
	var req model.RequestAdminAkk
	if err:= c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	admin, err := h.servise.CreateAdmin(req)
	if err!=nil{
		c.JSON(http.StatusConflict, gin.H{"error": "Admin name already exists"})
		return
	}
	c.JSON(http.StatusCreated, admin)
}