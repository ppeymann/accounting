package bank

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ppeymann/accounting.git"
	"github.com/ppeymann/accounting.git/server"
	"github.com/ppeymann/accounting.git/services"
)

type handler struct {
	service services.BankService
}

func NewHandler(service services.BankService, s *server.Server) services.BankHandler {
	handler := &handler{
		service: service,
	}
	group := s.Router.Group("/api/v1/bank")
	{
		group.Use(s.Authenticate())
		{
			group.POST("/create", handler.Create)
			group.GET("/all", handler.GetAllBank)
		}
	}
	return handler
}

// Create a bank account http request
//
// @BasePath			/api/v1/bank/create
// @Summary				create bank account
// @Description			create bank account with specified input
// @Tags				bank
// @Accept				json
// @Produce				json
//
// @Param				input	body		services.BankAccountInput	true	"account info"
// @Success				200		{object}	accounting.BaseResult "Always returns status 200 but body contains errors"
// @Router				/bank/create	[post]
// @Security			Authenticate bearer
func (h *handler) Create(ctx *gin.Context) {
	input := &services.BankAccountInput{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		ctx.JSON(http.StatusBadRequest, accounting.BaseResult{
			Errors: []string{accounting.ProvideRequiredJsonBody},
		})

		return
	}

	// call service
	result := h.service.Create(input, ctx)
	ctx.JSON(http.StatusOK, result)
}

// GetAllBank handles get all bank http request
//
// @BasePath			/api/v1/bank/all
// @Summary				get all bank
// @Description			get all bank information
// @Tags				bank
// @Accept				json
// @Produce				json
//
// @Success				200		{object}	accounting.BaseResult "Always returns status 200 but body contains errors"
// @Router				/bank/all		[get]
// @Security			Authenticate bearer
func (h *handler) GetAllBank(ctx *gin.Context) {
	// call service
	result := h.service.GetAllBank(ctx)
	ctx.JSON(http.StatusOK, result)
}
