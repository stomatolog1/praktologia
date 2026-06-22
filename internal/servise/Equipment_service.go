package servise

import (
	"github.com/stomatolog1/praktologia/internal/model"
	"github.com/stomatolog1/praktologia/internal/repository"
)

type EquipmentServise struct {
	repo *repository.EquipmentRepository
}

func NewEquipmentService(repo *repository.EquipmentRepository) *EquipmentServise {
	return &EquipmentServise{repo: repo}
}

func (s *EquipmentServise) CreateEquipment(req model.RequsetEquipment) (*model.Equipment, error) {
	equipment := &model.Equipment{
		Name: req.Name,
		Description: req.Description,
		TypeOperating: req.TypeOperating,
		RentalCost: req.RentalCost,
		PayTime: req.PayTime,
		Cost: req.Cost,
	}

	if err := s.repo.Create(equipment); err != nil {
		return nil, err
	}

	return equipment, nil
}

func (s *EquipmentServise) GetAll() ([]model.Equipment, error) {
	return s.repo.GetAll()
}

func (s *EquipmentServise) GetAllByEquipment(equipmentID uint) ([]model.Equipment, error) {
	return s.repo.GetAllByEquipment(equipmentID)
}