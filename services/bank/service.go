package bank

import (
	"net/http"

	"github.com/gin-gonic/gin"
	accounting "github.com/ppeymann/accounting.git"
	"github.com/ppeymann/accounting.git/auth"
	"github.com/ppeymann/accounting.git/services"
	"github.com/ppeymann/accounting.git/utils"
)

type service struct {
	repo services.BankRepository
}

func NewService(repo services.BankRepository) services.BankService {
	return &service{
		repo: repo,
	}
}

// Create implements services.BankService.
func (s *service) Create(input *services.BankAccountInput, ctx *gin.Context) *accounting.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{accounting.AuthorizationFailed},
		}
	}

	bank, err := s.repo.Create(input, claims.Subject)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	return &accounting.BaseResult{
		Status: http.StatusOK,
		Result: bank,
	}
}