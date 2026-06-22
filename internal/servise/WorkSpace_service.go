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
		AdminID:     req.AdminID,
		Name:        req.Name,
		Discription: req.Discription,
		TotalCost:   req.TotalCost,
		Resources:   req.Resources,
	}

	if err := s.repo.Create(WorkSpace); err != nil {
		return nil, err
	}

	return WorkSpace, nil
}

func (s *WorkSpaceServise) GetAll() ([]model.WorkSpace, error) {
	return s.repo.GetAll()
}

func (s *WorkSpaceServise) GetByAdminID(adminID uint) ([]model.WorkSpace, error) {
	return s.repo.GetByAdminID(adminID)
}

func (s *WorkSpaceServise) UpdateWorkSpaceStatus(id uint, status string) (*model.WorkSpace, error) {
    return s.repo.UpdateStatus(id, status)
}

func (s *WorkSpaceServise) DeleteWorkSpace(id uint) error {
    // potenteally delete projects in this workspace
    return s.repo.Delete(id)
}

func (s *WorkSpaceServise) GetAllByWorkSpace(SotrudnikID uint) ([]model.WorkSpace, error) {
	return s.repo.GetAllByWorkSpace(SotrudnikID)
}

func (s *WorkSpaceServise)  GetProjectsByAdmin(WorkSpaceID []uint) ([]model.WorkSpace, error){
	return s.repo.GetProjectsByAdmin(WorkSpaceID)
}