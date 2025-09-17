package routes

import (
	"user-service/controllers"
	"user-service/middlewares"

	"github.com/gin-gonic/gin"
)

type UserRoute struct {
	controller controllers.IRegistryController
	group      *gin.RouterGroup
}

type IUserRoute interface {
	Run()
}

func NewUserRoute(controller controllers.IRegistryController, group *gin.RouterGroup) IUserRoute {
	return &UserRoute{
		controller: controller,
		group:      group,
	}
}

func (r *UserRoute) Run() {
	group := r.group.Group("/auth")
	group.GET("/user", middlewares.Authenticate(), r.controller.GetUserController().GetUserLogin)
	group.GET("/:uuid", middlewares.Authenticate(), r.controller.GetUserController().GetUserByUUID)
	group.POST("/login", r.controller.GetUserController().Login)
	group.POST("/register", r.controller.GetUserController().Register)
	group.PUT("/:uuid", middlewares.Authenticate(), r.controller.GetUserController().Update)
}
