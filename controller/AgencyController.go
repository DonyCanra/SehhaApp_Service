package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sehha/dto/agency"
	"sehha/helper"
	"sehha/service"
	"sehha/util"
)

type AgencyController interface {
	Save(ctx *gin.Context)
}

type agencyController struct {
	agencyService service.AgencyService
}

func (a agencyController) Save(ctx *gin.Context) {
	//TODO implement me
	var response helper.Response
	var requestAgency agency.RequestAgencyDTO
	if err := ctx.ShouldBindJSON(&requestAgency); err != nil {
		var data = util.ValidationBindJsonField(err)
		ctx.JSON(http.StatusBadRequest, data)
		return
	} else {
		result, err, status := a.agencyService.SaveAndUpdate(requestAgency)
		if err != nil {
			response = helper.ResponseError(err.Error(), status)
			ctx.JSON(status, response)
		} else {
			response = helper.ResponseSuccess(result)
			ctx.JSON(status, response)
		}
	}
}

func NewAgencyController(agencyService service.AgencyService) AgencyController {
	return &agencyController{
		agencyService: agencyService,
	}
}
