package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ppeymann/accounting.git"
	"github.com/ppeymann/accounting.git/server"
	"github.com/ppeymann/accounting.git/services"
)

type handler struct {
	service services.AccountService
}

func NewHandler(svc services.AccountService, s *server.Server) services.AccountHandler {
	handler := &handler{
		service: svc,
	}

	group := s.Router.Group("/api/v1/account")
	{
		group.POST("/signup", handler.SignUp)
		group.POST("/signin", handler.SignIn)
		group.Use(s.Authenticate())
		{
			group.PATCH("/change_name", handler.ChangeName)
			group.PATCH("/change_currency", handler.ChangeCurrency)

		}
	}

	return handler
}

// SignUp handles signing up http request.
//
// @BasePath 		/api/v1/account/signup
// @Summary			signing up a new account
// @Description 	create new account with specified mobile and expected info
// @Tags 			account
// @Accept 			json
// @Produce 		json
//
// @Param			input		body		services.LoginInputDTO	true	"account info"
// @Success 		200 		{object} 	services.TokenBundleOutput		"always returns status 200 but body contains errors"
// @Router 			/account/signup	[post]
func (h *handler) SignUp(ctx *gin.Context) {
	input := &services.LoginInputDTO{}

	// get input from Body
	if err := ctx.ShouldBindJSON(input); err != nil {
		ctx.JSON(http.StatusBadRequest, accounting.BaseResult{
			Errors: []string{accounting.ProvideRequiredJsonBody},
		})

		return
	}

	// call associated service method for expected request
	result := h.service.SignUp(input, ctx)
	ctx.JSON(result.Status, result)

}

// SignIn handles sign in to application http request
//
// @BasePath			/api/v1/account/signin
// @Summary				signing in to account
// @Description			signing in to account that already signed up
// @Tags				account
// @Accept				json
// @Produce				json
//
// @Param				input	body	services.LoginInputDTO	true	"account info"
// @Success				200		{object}	services.TokenBundleOutput	"always returns status 200 but body contains errors"
// @Router				/account/signin	[post]
func (h *handler) SignIn(ctx *gin.Context) {
	input := &services.LoginInputDTO{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		ctx.JSON(http.StatusBadRequest, accounting.BaseResult{
			Errors: []string{accounting.ProvideRequiredJsonBody},
		})

		return
	}

	// call associated service method expected request
	result := h.service.SignIn(input, ctx)
	ctx.JSON(result.Status, result)
}

// ChangeName handles change name http request
//
// @BasePath			/api/v1/account/change_name
// @Summary				change name
// @Description			change name of account
// @Tags				account
// @Accept				json
// @Produce				json
//
// @Param				input	body	services.NameInput	true	"name that is for change"
// @Success				200		{object}	services.AccountEntity	"always returns status 200 but body contains errors"
// @Router				/account/change_name	[patch]
func (h *handler) ChangeName(ctx *gin.Context) {
	input := &services.NameInput{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		ctx.JSON(http.StatusBadRequest, accounting.BaseResult{
			Errors: []string{accounting.ProvideRequiredJsonBody},
		})

		return
	}

	// call associated service method expected request
	result := h.service.ChangeName(input, ctx)
	ctx.JSON(result.Status, result)
}

// ChangeCurrency handles change currency http request
//
// @BasePath			/api/v1/account/change_currency
// @Summary				change currency
// @Description			change currency of account
// @Tags				account
// @Accept				json
// @Produce				json
//
// @Param				input	body	services.ChangeCurrencyInput	true	"currency that is for change"
// @Success				200		{object}	services.AccountEntity	"always returns status 200 but body contains errors"
// @Router				/account/change_currency	[patch]
func (h *handler) ChangeCurrency(ctx *gin.Context) {
	input := &services.ChangeCurrencyInput{}

	if err := ctx.ShouldBindJSON(input); err != nil {
		ctx.JSON(http.StatusBadRequest, accounting.BaseResult{
			Errors: []string{accounting.ProvideRequiredJsonBody},
		})

		return
	}

	// call associated service method expected request
	result := h.service.ChangeCurrency(input, ctx)
	ctx.JSON(result.Status, result)
}
