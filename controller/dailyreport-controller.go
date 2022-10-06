package controller

import (
	"dailyreport/helper"
	"dailyreport/models/web"
	"dailyreport/service"
	"fmt"
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

func (drc *dailyreportController) All(context *gin.Context) {
	dreports := drc.dailyreportService.All()
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   dreports,
	}
	context.JSON(http.StatusOK, webResponse)
}

func (drc *dailyreportController) FindById(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	dreport, err := drc.dailyreportService.FindById(uint(id))

	const (
		layoutISO = "2006-01-02"
		layoutUS  = "January 2, 2006"
	)
	date := dreport.CreatedAt
	// t, _ := time.Parse(layoutISO, date)

	// dreport.StartTime.Format(time.Kitchen)
	// dreport.EndTime.Format(time.Kitchen)

	t1 := dreport.StartTime
	t2 := dreport.EndTime

	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Printf("Daily reported task took %v of work.\n", t2.Sub(t1))

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
		Data:     dreport,
		Date:     date.Format(layoutUS),
		Duration: diff.String(),
	}
	context.JSON(http.StatusOK, webResponse)
}

func (drc *dailyreportController) Insert(context *gin.Context) {
	var request web.DailyReportRequest

	err := context.BindJSON(&request)
	ok := helper.ValidationError(context, err)
	if ok {
		return
	}

	request.User_id = 1

	dreport, err := drc.dailyreportService.Create(request)

	const (
		layoutISO = "2006-01-02"
		layoutUS  = "January 2, 2006"
	)
	date := dreport.CreatedAt

	t1 := dreport.StartTime
	t2 := dreport.EndTime
	fmt.Printf("Daily reported task took %v of work.\n", t2.Sub(t1))
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
		Data:     dreport,
		Date:     date.Format(layoutUS),
		Duration: diff.String(),
	}
	context.JSON(http.StatusOK, webResponse)
}

func (drc *dailyreportController) Update(context *gin.Context) {
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
	dreport, err := drc.dailyreportService.Update(request)

	const (
		layoutISO = "2006-01-02"
		layoutUS  = "January 2, 2006"
	)
	date := dreport.CreatedAt

	t1 := dreport.StartTime
	t2 := dreport.EndTime
	fmt.Printf("Daily reported task took %v of work.\n", t2.Sub(t1))
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
		Data:     dreport,
		Date:     date.Format(layoutUS),
		Duration: diff.String(),
	}
	context.JSON(http.StatusOK, webResponse)
}

func (drc *dailyreportController) Delete(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	err = drc.dailyreportService.Delete(uint(id))
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
