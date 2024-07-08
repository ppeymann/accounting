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
// @Param				fromYear	query	int	true	"start year"
// @Param				fromMonth	query	int	true	"start month"
// @Param				toYear		query	int	true	"end year"
// @Param				toMonth		query	int	true	"end month"
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
