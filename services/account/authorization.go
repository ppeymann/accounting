package account

import (
	"github.com/gin-gonic/gin"
	"github.com/ppeymann/accounting.git"
	"github.com/ppeymann/accounting.git/services"
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
func (a *authorizationService) SignUp(input *services.LoginInputDTO, ctx *gin.Context) accounting.BaseResult {
	panic("unimplemented")
}
