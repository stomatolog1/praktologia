package repository

import (
	"github.com/stomatolog1/praktologia/internal/model"
	"gorm.io/gorm"
)

type ProjectRepository struct {
	db *gorm.DB
	projects []*model.Project
	counter uint
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{
		db: db,
		projects: make([]*model.Project, 0),
		counter: 0,
	}
}

func (r *ProjectRepository) Create(project *model.Project) error {
	r.counter++
	project.ID = r.counter
	r.projects = append(r.projects, project)
	return nil
}

func (r *ProjectRepository) GetAllByProject(ProjectID uint) ([]model.Project, error) {
	return []model.Project{}, nil
}

func (r *ProjectRepository) GetProjectsByWorkSpace(projectID []uint) ([]model.Project, error) {
	return []model.Project{}, nil
}
