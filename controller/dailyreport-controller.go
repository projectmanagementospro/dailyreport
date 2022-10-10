package controller

import (
	"dailyreport/helper"
	"dailyreport/models/web"
	"dailyreport/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DailyReportController interface {
	All(context *gin.Context)
	FindById(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type dailyreportController struct {
	dailyreportService service.DailyReportService
}

func NewDailyReportController(dailyreportService service.DailyReportService) DailyReportController {
	return &dailyreportController{
		dailyreportService: dailyreportService,
	}
}

func (dReportController *dailyreportController) All(context *gin.Context) {
	dreports := dReportController.dailyreportService.All()
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   dreports,
	}
	context.JSON(http.StatusOK, webResponse)
}

func (dReportController *dailyreportController) FindById(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	dreport, err := dReportController.dailyreportService.FindById(uint(id))

	const (
		layoutISO = "2006-01-02"
		layoutUS  = "January 2, 2006"
	)
	date := dreport.CreatedAt
	// t, _ := time.Parse(layoutISO, date)

	ok = helper.NotFoundError(context, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:     http.StatusOK,
		Status:   "Success",
		Errors:   "",
		Data:     dreport,
		Date:     date.Format(layoutUS),
		Duration: "",
	}
	context.JSON(http.StatusOK, webResponse)
}

func (dReportController *dailyreportController) Insert(context *gin.Context) {
	var request web.DailyReportRequest

	err := context.BindJSON(&request)
	ok := helper.ValidationError(context, err)
	if ok {
		return
	}

	// request.UserId = 1

	dreport, err := dReportController.dailyreportService.Create(request)

	const (
		layoutISO = "2006-01-02"
		layoutUS  = "January 2, 2006"
	)
	date := dreport.CreatedAt

	ok = helper.InternalServerError(context, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:     http.StatusOK,
		Status:   "Success",
		Errors:   "",
		Data:     dreport,
		Date:     date.Format(layoutUS),
		Duration: "",
	}
	context.JSON(http.StatusOK, webResponse)
}

func (dReportController *dailyreportController) Update(context *gin.Context) {
	var request web.DailyReportUpdateRequest
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	request.ID = uint(id)
	err = context.BindJSON(&request)
	ok = helper.ValidationError(context, err)
	if ok {
		return
	}
	dreport, err := dReportController.dailyreportService.Update(request)

	const (
		layoutISO = "2006-01-02"
		layoutUS  = "January 2, 2006"
	)
	date := dreport.CreatedAt

	ok = helper.InternalServerError(context, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:     http.StatusOK,
		Status:   "Success",
		Errors:   "",
		Data:     dreport,
		Date:     date.Format(layoutUS),
		Duration: "",
	}
	context.JSON(http.StatusOK, webResponse)
}

func (dReportController *dailyreportController) Delete(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	err = dReportController.dailyreportService.Delete(uint(id))
	ok = helper.NotFoundError(context, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:     http.StatusOK,
		Status:   "Success",
		Errors:   "",
		Data:     "Daily report has been removed",
		Date:     "Outdated",
		Duration: "Vanished",
	}
	context.JSON(http.StatusOK, webResponse)
}
