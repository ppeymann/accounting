package expenses

import (
	"net/http"

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
