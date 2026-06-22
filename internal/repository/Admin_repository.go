package repository

import(
	"gorm.io/gorm"
	"github.com/stomatolog1/praktologia/internal/model"
)

type AdminAkkRepository struct{
	db *gorm.DB
	admins []*model.AdminAkk
	counter uint
}

func NewAdminAkkRepository(db *gorm.DB) *AdminAkkRepository{
	return &AdminAkkRepository{
		db: db,
		admins: make([]*model.AdminAkk, 0),
		counter: 0,
	}
}

func (r *AdminAkkRepository) Create(admin *model.AdminAkk) error {
	r.counter++
	admin.ID = r.counter
	r.admins = append(r.admins, admin)
	return nil
}

func (r *AdminAkkRepository) FindByLogin(login string) (*model.AdminAkk, error) {
	for _, admin := range r.admins {
		if admin.Login == login {
			return admin, nil
		}
	}
	return nil, nil
}
