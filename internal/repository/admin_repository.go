package repository

import (
	"pg_pritani/backend/internal/domain"

	"gorm.io/gorm"
)

type AdminRepository interface {
	FindAll() ([]*domain.Admin, error)
	FindByID(id uint) (*domain.Admin, error)
	Update(admin *domain.Admin) error
	Delete(id uint) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db}
}

func (r *adminRepository) FindAll() ([]*domain.Admin, error) {
	var admins []*domain.Admin
	err := r.db.Find(&admins).Error
	return admins, err
}

func (r *adminRepository) FindByID(id uint) (*domain.Admin, error) {
	var admin domain.Admin
	err := r.db.Find(&admin, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &admin, err
}

func (r *adminRepository) Update(admin *domain.Admin) error {
	return r.db.Save(admin).Error
}

func (r *adminRepository) Delete(id uint) error {
	return r.db.Delete(id).Error
}
