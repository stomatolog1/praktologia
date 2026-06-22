package servise

import (
	"github.com/stomatolog1/praktologia/internal/model"
	"github.com/stomatolog1/praktologia/internal/repository"
)

type ExecutorServise struct {
	repo *repository.ExecutorRepository
}

func NewExecutorService(repo *repository.ExecutorRepository) *ExecutorServise {
	return &ExecutorServise{repo: repo}
}

func (s *ExecutorServise) CreateExecutor(req model.RequsetExecutor) (*model.Executor, error) {
	executor := &model.Executor{
		Name: req.Name,
		SecondName: req.SecondName,
		Surname: req.Surname,
		TypeDesign: req.TypeDesign,
		Tax: req.Tax,
		PayTime: req.PayTime,
		Cost: req.Cost,
	}

	if err := s.repo.Create(executor); err != nil {
		return nil, err
	}

	return executor, nil
}

func (s *ExecutorServise) GetAll() ([]model.Executor, error) {
	return s.repo.GetAll()
}

func (s *ExecutorServise) GetAllByExecutor(ExecutorID uint) ([]model.Executor, error) {
	return s.repo.GetAllByExecutor(ExecutorID)
}