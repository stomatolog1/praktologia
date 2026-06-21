package repository

import (
	"github.com/stomatolog1/praktologia/internal/model"
	"gorm.io/gorm"
)

type EquipmentRepository struct {
	db *gorm.DB
	equipment []*model.Equipment
	counter uint
}

func NewEquipmentRepository(db *gorm.DB) *EquipmentRepository {
	return &EquipmentRepository{
		db: db,
		equipment: make([]*model.Equipment, 0),
		counter: 0,
	}
}

func (r *EquipmentRepository) Create(equipment *model.Equipment) error {
	r.counter++
	equipment.ID = r.counter
	r.equipment = append(r.equipment, equipment)
	return nil
}

func (r *EquipmentRepository) GetAllByEquipment(equipmentID uint) ([]model.Equipment, error) {
	return []model.Equipment{}, nil
}
