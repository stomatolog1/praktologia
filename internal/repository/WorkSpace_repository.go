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

func (r *WorkSpaceRepository) GetAll() ([]model.WorkSpace, error) {
	result := make([]model.WorkSpace, len(r.workspaces))
	for i, w := range r.workspaces {
		result[i] = *w
	}
	return result, nil
}

func (r *WorkSpaceRepository) GetByAdminID(adminID uint) ([]model.WorkSpace, error) {
	result := make([]model.WorkSpace, 0)
	for _, w := range r.workspaces {
		if w.AdminID == adminID {
			result = append(result, *w)
		}
	}
	return result, nil
}

func (r *WorkSpaceRepository) GetAllByWorkSpace(WorkSpaceID uint) ([]model.WorkSpace, error) {
	return []model.WorkSpace{}, nil
}

func (r *WorkSpaceRepository) GetProjectsByAdmin(WorkSpaceID []uint) ([]model.WorkSpace, error) {
	return []model.WorkSpace{}, nil
}
