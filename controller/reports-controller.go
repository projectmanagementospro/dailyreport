package controller

import (
	"dailyreport/helper"
	"dailyreport/models/web"
	"dailyreport/service"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ReportsController interface {
	All(context *gin.Context)
	FindById(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type reportsController struct {
	reportsService service.ReportsService
}

func NewReportsController(reportsService service.ReportsService) ReportsController {
	return &reportsController{
		reportsService: reportsService,
	}
}

func (reportsController *reportsController) All(context *gin.Context) {
	reports := reportsController.reportsService.All()
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   reports,
	}
	context.JSON(http.StatusOK, webResponse)
}

func (reportsController *reportsController) FindById(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	report, err := reportsController.reportsService.FindById(uint(id))

	const (
		layoutISO = "2006-01-02"
		layoutUS  = "January 2, 2006"
	)
	date := report.CreatedAt
	// t, _ := time.Parse(layoutISO, date)

	report.StartTime.Format(time.Kitchen)
	report.EndTime.Format(time.Kitchen)

	t1 := report.StartTime
	t2 := report.EndTime

	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Printf("Reported task took %v of work.\n", t2.Sub(t1))

	diff := t2.Sub(t1)

	// out := time.Time{}.Add(diff)
	// fmt.Println(out.Format("15:04:05"))

	ok = helper.NotFoundError(context, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:     http.StatusOK,
		Status:   "Success",
		Errors:   "",
		Data:     report,
		Date:     date.Format(layoutUS),
		Duration: diff.String(),
	}
	context.JSON(http.StatusOK, webResponse)
}

func (reportsController *reportsController) Insert(context *gin.Context) {
	var request web.ReportsRequest

	err := context.BindJSON(&request)
	ok := helper.ValidationError(context, err)
	if ok {
		return
	}

	// request.UserId = 1

	report, err := reportsController.reportsService.Create(request)

	const (
		layoutISO = "2006-01-02"
		layoutUS  = "January 2, 2006"
	)
	date := report.CreatedAt

	t1 := report.StartTime
	t2 := report.EndTime
	fmt.Printf("Reported task took %v of work.\n", t2.Sub(t1))
	diff := t2.Sub(t1)

	// out := time.Time{}.Add(diff)
	// fmt.Println(out.Format("15:04:05"))

	ok = helper.InternalServerError(context, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:     http.StatusOK,
		Status:   "Success",
		Errors:   "",
		Data:     report,
		Date:     date.Format(layoutUS),
		Duration: diff.String(),
	}
	context.JSON(http.StatusOK, webResponse)
}

func (reportsController *reportsController) Update(context *gin.Context) {
	var request web.ReportsUpdateRequest
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
	report, err := reportsController.reportsService.Update(request)

	const (
		layoutISO = "2006-01-02"
		layoutUS  = "January 2, 2006"
	)
	date := report.CreatedAt

	t1 := report.StartTime
	t2 := report.EndTime
	fmt.Printf("Reported task took %v of work.\n", t2.Sub(t1))
	diff := t2.Sub(t1)

	// out := time.Time{}.Add(diff)

	ok = helper.InternalServerError(context, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:     http.StatusOK,
		Status:   "Success",
		Errors:   "",
		Data:     report,
		Date:     date.Format(layoutUS),
		Duration: diff.String(),
	}
	context.JSON(http.StatusOK, webResponse)
}

func (reportsController *reportsController) Delete(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	err = reportsController.reportsService.Delete(uint(id))
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
