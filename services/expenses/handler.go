package expenses

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ppeymann/accounting.git"
	"github.com/ppeymann/accounting.git/server"
	"github.com/ppeymann/accounting.git/services"
)

type handler struct {
	service services.ExpensesService
}

func NewHandler(service services.ExpensesService, s *server.Server) services.ExpensesHandler {
	handler := &handler{
		service: service,
	}

	group := s.Router.Group("/api/v1/expenses")
	{
		group.Use(s.Authenticate())
		{
			group.POST("/create", handler.Create)
			group.GET("/get_all", handler.GetAll)
			group.GET("/get_period_time", handler.GetPeriodTime)
			group.GET("/get_in_month/:year/:month", handler.GetInMonth)
			group.DELETE("/:id", handler.DeleteExpenses)
			group.PUT("/:id", handler.UpdateExpenses)
			group.GET("/:id", handler.GetByID)
		}
	}

	return handler
}

// Create expenses http request
//
// @BasePath			/api/v1/expenses/create
// @Summary				create expenses
// @Description			create expenses with expenses input
// @Tags				expenses
// @Accept				json
// @Produce				json
//
// @Param				input	body	services.ExpensesInput	true	"expenses input"
// @Success				200		{object}	accounting.BaseResult{result=services.ExpensesEntity}		"always returns status 200 but body contains errors"
// @Router				/expenses/create	[post]
// @Security			Authenticate Bearer
func (h *handler) Create(ctx *gin.Context) {
	input := &services.ExpensesInput{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		ctx.JSON(http.StatusBadRequest, accounting.BaseResult{
			Errors: []string{accounting.ProvideRequiredJsonBody},
		})

		return
	}

	result := h.service.Create(input, ctx)
	ctx.JSON(result.Status, result)
}

// GetAll is for get all expenses http request
//
// @BasePath			/api/v1/expenses/get_all
// @Summary				get all expenses
// @Description			get all expenses with specified id
// @Tags				expenses
// @Accept				json
// @Produce				json
//
// @Success				200		{object}	accounting.BaseResult{result=[]services.ExpensesEntity}	"always returns status 200 but body contains errors"
// @Router				/expenses/get_all	[get]
// @Security			Authenticate Bearer
func (h *handler) GetAll(ctx *gin.Context) {

	// call service
	result := h.service.GetAll(ctx)
	ctx.JSON(result.Status, result)
}

// GetPeriodTime is handler for get expenses from specific date to specific date
//
// @BasePath			/api/v1/expenses/get_period_time
// @Summary				get expenses
// @Description			get expenses in period time
// @Tags				expenses
// @Accept				json
// @Produce				json
//
// @Param				fromYear	query	string	true	"start year"
// @Param				fromMonth	query	string	true	"start month"
// @Param				toYear		query	string	true	"end year"
// @Param				toMonth		query	string	true	"end month"
// @Success				200		{object}	accounting.BaseResult{result=[]services.ExpensesEntity}	"always returns status 200 but body contains errors"
// @Router				/expenses/get_period_time	[get]
// @Security			Authenticate Bearer
func (h *handler) GetPeriodTime(ctx *gin.Context) {
	fromYearStr := ctx.Query("fromYear")
	fromMonthStr := ctx.Query("fromMonth")
	toYearStr := ctx.Query("toYear")
	toMonthStr := ctx.Query("toMonth")

	fromYear, err := strconv.Atoi(fromYearStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, accounting.BaseResult{
			Errors: []string{"Invalid fromYear parameter"},
		})
		return
	}

	fromMonth, err := strconv.Atoi(fromMonthStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, accounting.BaseResult{
			Errors: []string{"Invalid fromMonth parameter"},
		})
		return
	}

	toYear, err := strconv.Atoi(toYearStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, accounting.BaseResult{
			Errors: []string{"Invalid toYear parameter"},
		})
		return
	}

	toMonth, err := strconv.Atoi(toMonthStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, accounting.BaseResult{
			Errors: []string{"Invalid toMonth parameter"},
		})
		return
	}

	input := &services.PeriodTimeInput{
		From: services.PeriodDateAndTime{
			Year:  fromYear,
			Month: fromMonth,
		},
		To: services.PeriodDateAndTime{
			Year:  toYear,
			Month: toMonth,
		},
	}

	result := h.service.GetPeriodTime(input, ctx)
	ctx.JSON(result.Status, result)
}

// GetInMonth is for get expenses in the specified month
//
// @BasePath			/api/v1/expenses/get_in_month
// @Summary				get expenses
// @Description			get expenses in the specified month
// @Tags				expenses
// @Accept				json
// @Produce				json
//
// @Param				year	path	string	true	"year"
// @Param				month	path	string	true	"month"
// @Success				200		{object}	accounting.BaseResult{result=[]services.ExpensesEntity}	"always returns status 200 but body contains errors"
// @Router				/expenses/get_in_month/{year}/{month}	[get]
func (h *handler) GetInMonth(ctx *gin.Context) {
	year, err := server.GetInt64Path("year", ctx)
	if err != nil {
		return
	}

	month, err := server.GetInt64Path("month", ctx)
	if err != nil {
		return
	}

	result := h.service.GetInMonth(int(year), int(month), ctx)
	ctx.JSON(result.Status, result)
}

// DeleteExpenses is for delete expenses
//
// @BasePath			/api/v1/expenses
// @Summary				delete expenses
// @Description			delete expenses
// @Tags				expenses
// @Accept				json
// @Produce				json
//
// @Param				id	path		string	true	"expenses id"
// @Success				200	{object}	accounting.BaseResult{result=int}	"always returns status 200 but body contains errors"
// @Router				/expenses/{id}	[delete]
// @Security			Authenticate Bearer
func (h *handler) DeleteExpenses(ctx *gin.Context) {
	id, err := server.GetPathUint64(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, accounting.BaseResult{
			Errors: []string{err.Error()},
		})

		return
	}

	result := h.service.DeleteExpenses(uint(id), ctx)
	ctx.JSON(result.Status, result)
}

// UpdateExpenses is for update expenses
//
// @BasePath			/api/v1/expenses
// @Summary				update expenses
// @Description			update expenses
// @Tags				expenses
// @Accept				json
// @Produce				json
//
// @Param				id		path		string	true	"expenses id"
// @Param				input	body		services.ExpensesInput	true	"expenses input"
// @Success				200		{object}	accounting.BaseResult{result=services.ExpensesEntity}	"always returns status 200 but body contains errors"
// @Router				/expenses/{id}	[put]
// @Security			Authenticate Bearer
func (h *handler) UpdateExpenses(ctx *gin.Context) {
	id, err := server.GetPathUint64(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, accounting.BaseResult{
			Errors: []string{err.Error()},
		})

		return
	}

	input := &services.ExpensesInput{}

	if err = ctx.ShouldBindJSON(input); err != nil {
		ctx.JSON(http.StatusBadRequest, accounting.BaseResult{
			Errors: []string{accounting.ProvideRequiredJsonBody},
		})

		return
	}

	result := h.service.UpdateExpenses(uint(id), input, ctx)
	ctx.JSON(result.Status, result)
}

// GetByID is for get expenses by id
//
// @BasePath			/api/v1/expenses
// @Summary				get expenses
// @Description			get expenses by id
// @Tags				expenses
// @Accept				json
// @Produce				json
//
// @Param				id	path		string	true	"expenses id"
// @Success				200	{object}	accounting.BaseResult{result=services.ExpensesEntity}	"always returns status 200 but body contains errors"
// @Router				/expenses/{id}	[get]
// @Security			Authenticate Bearer
func (h *handler) GetByID(ctx *gin.Context) {
	id, err := server.GetPathUint64(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, accounting.BaseResult{
			Errors: []string{err.Error()},
		})

		return
	}

	result := h.service.GetByID(uint(id), ctx)
	ctx.JSON(result.Status, result)
}
