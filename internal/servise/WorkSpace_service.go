package servise

import (
	"github.com/stomatolog1/praktologia/internal/model"
	"github.com/stomatolog1/praktologia/internal/repository"
)

type WorkSpaceServise struct {
	repo *repository.WorkSpaceRepository
}

func NewWorkSpaceService(repo *repository.WorkSpaceRepository) *WorkSpaceServise {
	return &WorkSpaceServise{repo: repo}
}

func (s *WorkSpaceServise) CreateWorkSpace(req model.RequestWorkSpace) (*model.WorkSpace, error) {
	WorkSpace := &model.WorkSpace{
		AdminID: req.AdminID,
		Name: req.Name,
		Discription: req.Discription,
	}

	if err := s.repo.Create(WorkSpace); err != nil {
		return nil, err
	}

	return WorkSpace, nil
}

func (s *WorkSpaceServise) GetAllByWorkSpace(SotrudnikID uint) ([]model.WorkSpace, error) {
	return s.repo.GetAllByWorkSpace(SotrudnikID)
}

func (s *WorkSpaceServise)  GetProjectsByAdmin(WorkSpaceID []uint) ([]model.WorkSpace, error){
	return s.repo.GetProjectsByAdmin(WorkSpaceID)
}