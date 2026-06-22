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

func (s *AdminServise) CreateAdmin(req model.RequestAdminAkk) (*model.AdminAkk, error) {
	existing, err := s.repo.FindByLogin(req.Login)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, errLoginExists
	}

	admin := &model.AdminAkk{
		Login:    req.Login,
		HashPass: HashPass(req.HashPass),
	}

	if err := s.repo.Create(admin); err != nil {
		return nil, err
	}

	return admin, nil
}

func (s *AdminServise) Login(req model.RequestAdminLogin) (*model.AdminResponse, error) {
	admin, err := s.repo.FindByLogin(req.Login)
	if err != nil {
		return nil, err
	}
	if admin == nil || admin.HashPass != HashPass(req.HashPass) {
		return nil, errInvalidCredentials
	}
	return &model.AdminResponse{ID: admin.ID, Login: admin.Login}, nil
}

var errLoginExists = &loginError{"login already exists"}
var errInvalidCredentials = &loginError{"invalid login or password"}

type loginError struct{ msg string }

func (e *loginError) Error() string { return e.msg }

func HashPass(p string) string{
	return p
}