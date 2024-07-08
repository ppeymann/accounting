package expenses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	accounting "github.com/ppeymann/accounting.git"
	"github.com/ppeymann/accounting.git/auth"
	"github.com/ppeymann/accounting.git/services"
	"github.com/ppeymann/accounting.git/utils"
)

type authorizationService struct {
	next services.ExpensesService
}

func NewAuthorizationService(service services.ExpensesService) services.ExpensesService {
	return &authorizationService{
		next: service,
	}
}

// Create implements services.ExpensesService.
func (a *authorizationService) Create(input *services.ExpensesInput, ctx *gin.Context) *accounting.BaseResult {
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

// GetAll implements services.ExpensesService.
func (a *authorizationService) GetAll(ctx *gin.Context) *accounting.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{accounting.AuthorizationFailed},
		}
	}

	return a.next.GetAll(ctx)
}
