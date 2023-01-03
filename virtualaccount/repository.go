package virtualaccount

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Kontrakva, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Kontrakva, error) {
	var kontrak []Kontrakva
	err := r.db.Find(&kontrak).Error
	return kontrak, err
}
