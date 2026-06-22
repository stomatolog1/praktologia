package servise

import (
	"github.com/stomatolog1/praktologia/internal/model"
	"github.com/stomatolog1/praktologia/internal/repository"
)

type NMAService struct {
	nmaRepo *repository.NMARepository
}

func NewNMAService(nmaRepo *repository.NMARepository) *NMAService {
	return &NMAService{nmaRepo: nmaRepo}
}

func (s *NMAService) GetNMAData(projectID uint) (*model.NMA, error) {
	return s.nmaRepo.GetNMAData(projectID)
}
