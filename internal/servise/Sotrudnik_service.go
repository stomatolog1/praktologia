package servise

import (
	"github.com/stomatolog1/praktologia/internal/model"
	"github.com/stomatolog1/praktologia/internal/repository"
)

type SotrudnikServise struct {
	repo *repository.SotrudnikRepository
}

func NewSotrudnikService(repo *repository.SotrudnikRepository) *SotrudnikServise {
	return &SotrudnikServise{repo: repo}
}

func (s *SotrudnikServise) CreateSotrudnik(req model.RequestSotrudnik) (*model.Sotrudnik, error) {
	sotrudnik := &model.Sotrudnik{
		Name: req.Name,
		SecondName: req.SecondName,
		Surname: req.Surname,
		Position: req.Position,
		PayMount: req.PayMount,
		Tax: req.Tax,
	}

	if err := s.repo.Create(sotrudnik); err != nil {
		return nil, err
	}

	return sotrudnik, nil
}

func (s *SotrudnikServise) GetAll() ([]model.Sotrudnik, error) {
	return s.repo.GetAll()
}

func (s *SotrudnikServise) DeleteSotrudnik(id uint) error {
	return s.repo.Delete(id)
}

func (s *SotrudnikServise) GetAllBySotrudnik(SotrudnikID uint) ([]model.Sotrudnik, error) {
	return s.repo.GetAllBySotrudnik(SotrudnikID)
}
