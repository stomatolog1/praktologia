package repository

import (
	"github.com/stomatolog1/praktologia/internal/model"
	"gorm.io/gorm"
)

type ExecutorRepository struct {
	db *gorm.DB
	executors []*model.Executor
	counter uint
}

func NewExecutorRepository(db *gorm.DB) *ExecutorRepository {
	return &ExecutorRepository{
		db: db,
		executors: make([]*model.Executor, 0),
		counter: 0,
	}
}

func (r *ExecutorRepository) Create(executor *model.Executor) error {
	r.counter++
	executor.ID = r.counter
	r.executors = append(r.executors, executor)
	return nil
}

func (r *ExecutorRepository) GetAll() ([]model.Executor, error) {
	result := make([]model.Executor, len(r.executors))
	for i, e := range r.executors {
		result[i] = *e
	}
	return result, nil
}

func (r *ExecutorRepository) Delete(id uint) error {
	for i, e := range r.executors {
		if e.ID == id {
			r.executors = append(r.executors[:i], r.executors[i+1:]...)
			return nil
		}
	}
	return gorm.ErrRecordNotFound
}

func (r *ExecutorRepository) GetAllByExecutor(ExecutorID uint) ([]model.Executor, error) {
	return []model.Executor{}, nil
}
