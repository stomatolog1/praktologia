package servise

import (
	"github.com/stomatolog1/praktologia/internal/model"
	"github.com/stomatolog1/praktologia/internal/repository"
)

type CommercialOfferService struct {
	offerRepo *repository.CommercialOfferRepository
}

func NewCommercialOfferService(offerRepo *repository.CommercialOfferRepository) *CommercialOfferService {
	return &CommercialOfferService{offerRepo: offerRepo}
}

func (s *CommercialOfferService) GetCommercialOfferData(projectID uint, customerID uint) (*model.CommercialOffer, error) {
	return s.offerRepo.GetCommercialOfferData(projectID, customerID)
}
