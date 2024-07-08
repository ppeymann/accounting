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
			group.GET("/by_id/:id", handler.GetByID)
			group.DELETE("/:id", handler.DeleteBankAccount)
			group.PUT("/:id", handler.UpdateBankAccount)
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
// @Success				200		{object}	accounting.BaseResult{result=services.BankAccountEntity} "Always returns status 200 but body contains errors"
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
// @Success				200		{object}	accounting.BaseResult{result=[]services.BankEntity} "Always returns status 200 but body contains errors"
// @Router				/bank/all		[get]
// @Security			Authenticate bearer
func (h *handler) GetAllBank(ctx *gin.Context) {
	// call service
	result := h.service.GetAllBank(ctx)
	ctx.JSON(http.StatusOK, result)
}

// GetByID is for get bank by id
//
// @BasePath			/api/v1/by_id
// @Summary				get bank
// @Description			get bank by id
// @Tags				bank
// @Accept				json
// @Produce				json
//
// @Param				id	path		string	true	"bank id"
// @Success				200	{object}	accounting.BaseResult{result=services.BankAccountEntity}	"always returns status 200 but body contains errors"
// @Router				/by_id/{id}	[get]
// @Security			Authenticate bearer
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

// DeleteBankAccount is for delete bank account
//
// @BasePath			/api/v1/bank
// @Summary				delete bank
// @Description			delete bank by id
// @Tags				bank
// @Accept				json
// @Produce				json
//
// @Param				id	path		string	true	"bank id"
// @Success				200	{object}	accounting.BaseResult{result=int}	"always returns status 200 but body contains errors"
// @Router				/bank/{id}	[delete]
// @Security			Authenticate bearer
func (h *handler) DeleteBankAccount(ctx *gin.Context) {
	id, err := server.GetPathUint64(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, accounting.BaseResult{
			Errors: []string{err.Error()},
		})

		return
	}

	result := h.service.DeleteBankAccount(uint(id), ctx)
	ctx.JSON(result.Status, result)
}

// UpdateBankAccount is for update bank account
//
// @BasePath			/api/v1/bank
// @Summary				update bank
// @Description			update bank by id
// @Tags				bank
// @Accept				json
// @Produce				json
//
// @Param				id		path		string	true	"bank id"
// @Param				input	body		services.BankAccountInput	true	"account info"
// @Success				200		{object}	accounting.BaseResult{result=services.BankAccountEntity}	"always returns status 200 but body contains errors"
// @Router				/bank/{id}	[put]
// @Security			Authenticate bearer
func (h *handler) UpdateBankAccount(ctx *gin.Context) {
	id, err := server.GetPathUint64(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, accounting.BaseResult{
			Errors: []string{err.Error()},
		})

		return
	}

	input := &services.BankAccountInput{}
	if err = ctx.ShouldBindJSON(input); err != nil {
		ctx.JSON(http.StatusBadRequest, accounting.BaseResult{
			Errors: []string{accounting.ProvideRequiredJsonBody},
		})

		return
	}

	result := h.service.UpdateBankAccount(uint(id), input, ctx)
	ctx.JSON(result.Status, result)
}
