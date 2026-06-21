package servise

import (
	"github.com/stomatolog1/praktologia/internal/model"
	"github.com/stomatolog1/praktologia/internal/repository"
)

type AdminServise struct{
	repo *repository.AdminAkkRepository
}

func NewAdminAkkService(repo *repository.AdminAkkRepository) *AdminServise{
	return &AdminServise{repo: repo}
}

func (s *AdminServise) CreateAdmin(req model.RequestAdminAkk)(*model.AdminAkk, error){
	admin := &model.AdminAkk{
		Login: req.Login,
		HashPass: HashPass(req.HashPass),
	}

	if err := s.repo.Create(admin); err != nil {
		return nil, err
	}

	return admin, nil
}

func HashPass(p string) string{
	return p
}