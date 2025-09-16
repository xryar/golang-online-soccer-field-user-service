package repositories

import (
	repositories "user-service/repositories/user"

	"gorm.io/gorm"
)

type Registry struct {
	db *gorm.DB
}

type IRegistryRepository interface {
	GetUser() repositories.IUserRepository
}

func NewRegistryRepository(db *gorm.DB) IRegistryRepository {
	return &Registry{db: db}
}

func (r *Registry) GetUser() repositories.IUserRepository {
	return repositories.NewUserRepository(r.db)
}
