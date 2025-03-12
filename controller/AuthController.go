package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sehha/dto"
	"sehha/helper"
	"sehha/service"
	"sehha/util"
	"fmt"
)

type AuthController interface {
	Login(ctx *gin.Context)
	RegisterUser(ctx *gin.Context)
}

type authController struct {
	userService service.UserService
}

func (a authController) RegisterUser(ctx *gin.Context) {
	//TODO implement me
	var response helper.Response
	var reqRegister dto.ReqRegisterDTO
	if err := ctx.ShouldBindJSON(&reqRegister); err != nil {
		var data = util.ValidationBindJsonField(err)
		ctx.JSON(http.StatusBadRequest, data)
		return
	} else {
		result, err, status := a.userService.SaveUser(reqRegister)
		if err != nil {
			response = helper.ResponseError(err.Error(), status)
			ctx.JSON(status, response)
		} else {
			response = helper.ResponseSuccess(result)
			ctx.JSON(status, response)
		}
	}
}

func (a authController) Login(ctx *gin.Context) {
	//TODO implement me
	var response helper.Response
	fmt.Println("response", response);
	var reqLogin dto.ReqLoginDTO
	if err := ctx.ShouldBindJSON(&reqLogin); err != nil {
		var data = util.ValidationBindJsonField(err)
		ctx.JSON(http.StatusBadRequest, data)
		return
	} else {
		result, err, status := a.userService.ValidateUser(reqLogin)
		if err != nil {
			response = helper.ResponseError(err.Error(), status)
			ctx.JSON(status, response)
		} else {
			generateToken, err := util.GenerateToken(result.Email, result.Password)
			if err != nil {
				response = helper.ResponseError(err.Error(), status)
				ctx.JSON(status, response)
			}
			response = helper.ResponseSuccess(generateToken)
			ctx.JSON(status, response)
		}
	}
}

func NewAuthController(userService service.UserService) AuthController {
	return &authController{
		userService: userService,
	}
}
