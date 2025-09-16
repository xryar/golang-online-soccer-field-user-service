package services

import (
	"user-service/repositories"
	services "user-service/services/user"
)

type Registry struct {
	repository repositories.IRegistryRepository
}

type IRegistryService interface {
	GetUser() services.IUserService
}

func NewRegistryService(repository repositories.IRegistryRepository) IRegistryService {
	return &Registry{repository: repository}
}

func (r *Registry) GetUser() services.IUserService {
	return services.NewUserService(r.repository)
}
