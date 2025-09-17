package controllers

import (
	controllers "user-service/controllers/user"
	"user-service/services"
)

type Registry struct {
	service services.IRegistryService
}

type IRegistryController interface {
	GetUserController() controllers.IUserController
}

func NewRegistryController(service services.IRegistryService) IRegistryController {
	return &Registry{service: service}
}

func (r *Registry) GetUserController() controllers.IUserController {
	return controllers.NewUserController(r.service)
}
