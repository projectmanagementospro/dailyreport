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

func (c *dailyreportController) All(context *gin.Context) {
	dreports := c.dailyreportService.All()
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   dreports,
	}
	context.JSON(http.StatusOK, webResponse)
}

func (c *dailyreportController) FindById(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	dreport, err := c.dailyreportService.FindById(uint(id))

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

func (c *dailyreportController) Insert(context *gin.Context) {
	var u web.DailyReportRequest

	err := context.BindJSON(&u)
	ok := helper.ValidationError(context, err)
	if ok {
		return
	}

	u.User_id = 1

	dreport, err := c.dailyreportService.Create(u)

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

func (c *dailyreportController) Update(context *gin.Context) {
	var u web.DailyReportUpdateRequest
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	u.ID = uint(id)
	err = context.BindJSON(&u)
	ok = helper.ValidationError(context, err)
	if ok {
		return
	}
	dreport, err := c.dailyreportService.Update(u)

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

func (c *dailyreportController) Delete(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	err = c.dailyreportService.Delete(uint(id))
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
