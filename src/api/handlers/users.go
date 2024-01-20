package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/salmantaghooni/golang-car-web-api/src/api/dto"
	"github.com/salmantaghooni/golang-car-web-api/src/api/helper"
	"github.com/salmantaghooni/golang-car-web-api/src/config"
	"github.com/salmantaghooni/golang-car-web-api/src/services"
)

type UsersHandler struct {
	services *services.UserService
}

func NewUsersHandler(cfg *config.Config) *UsersHandler {
	service := services.NewUserService(cfg)
	return &UsersHandler{services: service}
}

func (u *UsersHandler) SendOtp(ctx *gin.Context) {
	req := new(dto.GetOtpRequest)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithError(nil, false, http.StatusBadRequest, err))
		return
	}
	err = u.services.SendOtp(req)
	if err != nil {
		errStatusCode := helper.TranslateErrorToStatusCode(err)
		ctx.AbortWithStatusJSON(errStatusCode, helper.GenerateBaseResponseWithError(nil, false, errStatusCode, err))
		return
	}

	//Call Send Internal SMS Service
	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, http.StatusCreated))
}
