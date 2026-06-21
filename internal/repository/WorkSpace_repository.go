package repository

import (
	"github.com/stomatolog1/praktologia/internal/model"
	"gorm.io/gorm"
)

type WorkSpaceRepository struct {
	db *gorm.DB
	workspaces []*model.WorkSpace
	counter uint
}

func NewWorkSpaceRepository(db *gorm.DB) *WorkSpaceRepository {
	return &WorkSpaceRepository{
		db: db,
		workspaces: make([]*model.WorkSpace, 0),
		counter: 0,
	}
}

func (r *WorkSpaceRepository) Create(WorkSpace *model.WorkSpace) error {
	r.counter++
	WorkSpace.ID = r.counter
	r.workspaces = append(r.workspaces, WorkSpace)
	return nil
}

func (r *WorkSpaceRepository) GetAllByWorkSpace(WorkSpaceID uint) ([]model.WorkSpace, error) {
	return []model.WorkSpace{}, nil
}

func (r *WorkSpaceRepository) GetProjectsByAdmin(WorkSpaceID []uint) ([]model.WorkSpace, error) {
	return []model.WorkSpace{}, nil
}
