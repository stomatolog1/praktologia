package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stomatolog1/praktologia/internal/servise"
)

type CommercialOfferHandler struct {
	offerService *servise.CommercialOfferService
}

func NewCommercialOfferHandler(offerService *servise.CommercialOfferService) *CommercialOfferHandler {
	return &CommercialOfferHandler{offerService: offerService}
}

func (h *CommercialOfferHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/commercial-offer/:projectID/:customerID", h.getCommercialOffer)
}

func (h *CommercialOfferHandler) getCommercialOffer(c *gin.Context) {
	projectID, err := strconv.ParseUint(c.Param("projectID"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	customerID, err := strconv.ParseUint(c.Param("customerID"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	offer, err := h.offerService.GetCommercialOfferData(uint(projectID), uint(customerID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, offer)
}
