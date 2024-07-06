package account

import (
	"github.com/gin-gonic/gin"
	"github.com/ppeymann/accounting.git"
	"github.com/ppeymann/accounting.git/auth"
	"github.com/ppeymann/accounting.git/services"
	"github.com/ppeymann/accounting.git/utils"
)

type authorizationService struct {
	next services.AccountService
}

func NewAuthorizationService(service services.AccountService) services.AccountService {
	return &authorizationService{
		next: service,
	}
}

// SignUp implements service.AccountService.
func (a *authorizationService) SignUp(input *services.LoginInputDTO, ctx *gin.Context) *accounting.BaseResult {
	return a.next.SignUp(input, ctx)
}

// SignIn implements services.AccountService.
func (a *authorizationService) SignIn(input *services.LoginInputDTO, ctx *gin.Context) *accounting.BaseResult {
	return a.next.SignIn(input, ctx)
}

// ChangeName implements services.AccountService.
func (a *authorizationService) ChangeName(input *services.NameInput, ctx *gin.Context) *accounting.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &accounting.BaseResult{
			Errors: []string{accounting.AuthorizationFailed},
		}
	}

	return a.next.ChangeName(input, ctx)
}
