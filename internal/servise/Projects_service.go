package servise

import (
	"github.com/stomatolog1/praktologia/internal/model"
	"github.com/stomatolog1/praktologia/internal/repository"
)

type ProjectServise struct {
	repo *repository.ProjectRepository
}

func NewProjectService(repo *repository.ProjectRepository) *ProjectServise {
	return &ProjectServise{repo: repo}
}

func (s *ProjectServise) CreateProject(req model.RequestProject) (*model.Project, error) {
	project := &model.Project{
		WorkSpaceID: req.WorkSpaceID,
		Name: req.Name,
		Discription: req.Discription,
		Deadline: req.Deadline,
		Price: req.Price,
		Workers: req.Workers,
		Equipments: req.Equipments,
	}

	if err := s.repo.Create(project); err != nil {
		return nil, err
	}

	return project, nil
}

func (s *ProjectServise) GetAll() ([]model.Project, error) {
	return s.repo.GetAll()
}

func (s *ProjectServise) GetByWorkSpace(workSpaceID uint) ([]model.Project, error) {
	return s.repo.GetByWorkSpace(workSpaceID)
}

func (s *ProjectServise) GetAllByProject(ProjectID uint) ([]model.Project, error) {
	return s.repo.GetAllByProject(ProjectID)
}

func (s *ProjectServise) GetProjectsByWorkSpace(ProjectID []uint) ([]model.Project, error){
	return s.repo.GetProjectsByWorkSpace(ProjectID)
}