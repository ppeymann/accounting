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
		group.GET("/signup", handler.SignUp)
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
