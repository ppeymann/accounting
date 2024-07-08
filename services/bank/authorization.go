package bank

import (
	"net/http"

	"github.com/gin-gonic/gin"
	accounting "github.com/ppeymann/accounting.git"
	"github.com/ppeymann/accounting.git/auth"
	"github.com/ppeymann/accounting.git/services"
	"github.com/ppeymann/accounting.git/utils"
)

type authorizationService struct {
	next services.BankService
}

func NewAuthorizationService(service services.BankService) services.BankService {
	return &authorizationService{
		next: service,
	}
}

// Create implements services.BankService.
func (a *authorizationService) Create(input *services.BankAccountInput, ctx *gin.Context) *accounting.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{accounting.AuthorizationFailed},
		}
	}

	return a.next.Create(input, ctx)
}

// GetAllBank implements services.BankService.
func (a *authorizationService) GetAllBank(ctx *gin.Context) *accounting.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{accounting.AuthorizationFailed},
		}
	}

	return a.next.GetAllBank(ctx)
}

// GetByID implements services.BankService.
func (a *authorizationService) GetByID(id uint, ctx *gin.Context) *accounting.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{accounting.AuthorizationFailed},
		}
	}

	return a.next.GetByID(id, ctx)
}
