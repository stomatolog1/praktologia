package repository

import (
	"github.com/stomatolog1/praktologia/internal/model"
	"gorm.io/gorm"
)

type SotrudnikRepository struct {
	db *gorm.DB
	sotrudniks []*model.Sotrudnik
	counter uint
}

func NewSotrudnikRepository(db *gorm.DB) *SotrudnikRepository {
	return &SotrudnikRepository{
		db: db,
		sotrudniks: make([]*model.Sotrudnik, 0),
		counter: 0,
	}
}

func (r *SotrudnikRepository) Create(sotrudnik *model.Sotrudnik) error {
	r.counter++
	sotrudnik.ID = r.counter
	r.sotrudniks = append(r.sotrudniks, sotrudnik)
	return nil
}

func (r *SotrudnikRepository) GetAll() ([]model.Sotrudnik, error) {
	result := make([]model.Sotrudnik, len(r.sotrudniks))
	for i, s := range r.sotrudniks {
		result[i] = *s
	}
	return result, nil
}

func (r *SotrudnikRepository) Delete(id uint) error {
	for i, s := range r.sotrudniks {
		if s.ID == id {
			r.sotrudniks = append(r.sotrudniks[:i], r.sotrudniks[i+1:]...)
			return nil
		}
	}
	return gorm.ErrRecordNotFound
}

func (r *SotrudnikRepository) GetAllBySotrudnik(SotrudnikID uint) ([]model.Sotrudnik, error) {
	return []model.Sotrudnik{}, nil
}
