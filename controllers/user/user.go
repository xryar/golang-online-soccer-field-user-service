package controllers

import (
	"net/http"
	errWrap "user-service/common/error"
	"user-service/common/response"
	"user-service/domain/dto"
	"user-service/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	service services.IRegistryService
}

type IUserController interface {
	Login(*gin.Context)
	Register(*gin.Context)
	Update(*gin.Context)
	GetUserLogin(*gin.Context)
	GetUserByUUID(*gin.Context)
}

func NewUserController(service services.IRegistryService) IUserController {
	return &UserController{service: service}
}

func (uc *UserController) Login(ctx *gin.Context) {
	request := &dto.LoginRequest{}

	err := ctx.ShouldBindJSON(request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResponse{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})

		return
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := errWrap.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: &errMessage,
			Data:    errResponse,
			Err:     err,
			Gin:     ctx,
		})

		return
	}

	user, err := uc.service.GetUser().Login(ctx, request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResponse{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})

		return
	}

	response.HttpResponse(response.ParamHTTPResponse{
		Code:  http.StatusOK,
		Data:  user.User,
		Token: &user.Token,
		Gin:   ctx,
	})
}

func (uc *UserController) Register(ctx *gin.Context) {
	request := &dto.RegisterRequest{}

	err := ctx.ShouldBindJSON(request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResponse{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})

		return
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := errWrap.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: &errMessage,
			Data:    errResponse,
			Err:     err,
			Gin:     ctx,
		})

		return
	}

	user, err := uc.service.GetUser().Register(ctx, request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResponse{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})

		return
	}

	response.HttpResponse(response.ParamHTTPResponse{
		Code: http.StatusOK,
		Data: user.User,
		Gin:  ctx,
	})
}

func (uc *UserController) Update(ctx *gin.Context) {
	request := &dto.UpdateRequest{}
	uuid := ctx.Param("uuid")

	err := ctx.ShouldBindJSON(request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResponse{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})

		return
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := errWrap.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: &errMessage,
			Data:    errResponse,
			Err:     err,
			Gin:     ctx,
		})

		return
	}

	user, err := uc.service.GetUser().Update(ctx, request, uuid)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResponse{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})

		return
	}

	response.HttpResponse(response.ParamHTTPResponse{
		Code: http.StatusOK,
		Data: user,
		Gin:  ctx,
	})
}

func (uc *UserController) GetUserLogin(ctx *gin.Context) {
	user, err := uc.service.GetUser().GetUserLogin(ctx.Request.Context())
	if err != nil {
		response.HttpResponse(response.ParamHTTPResponse{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResponse{
		Code: http.StatusOK,
		Data: user,
		Gin:  ctx,
	})
}

func (uc *UserController) GetUserByUUID(ctx *gin.Context) {
	user, err := uc.service.GetUser().GetUserByUUID(ctx.Request.Context(), ctx.Param("uuid"))
	if err != nil {
		response.HttpResponse(response.ParamHTTPResponse{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResponse{
		Code: http.StatusOK,
		Data: user,
		Gin:  ctx,
	})
}
