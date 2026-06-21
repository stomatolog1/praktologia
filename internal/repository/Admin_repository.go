package repository

import(
	"gorm.io/gorm"

	"github.com/stomatolog1/praktologia/internal/model"
)

type AdminAkkRepository struct{
	db *gorm.DB
}

func NewAdminAkkRepository(db *gorm.DB) *AdminAkkRepository{
	return &AdminAkkRepository{db: db}
}

func (r *AdminAkkRepository) Create(admin *model.AdminAkk) error{
	return r.db.Create(user).Error
}