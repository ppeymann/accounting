package expenses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	accounting "github.com/ppeymann/accounting.git"
	"github.com/ppeymann/accounting.git/auth"
	"github.com/ppeymann/accounting.git/services"
	"github.com/ppeymann/accounting.git/utils"
)

type service struct {
	repo services.ExpensesRepository
}

func NewService(repo services.ExpensesRepository) services.ExpensesService {
	return &service{
		repo: repo,
	}
}

// Create implements services.ExpensesService.
func (s *service) Create(input *services.ExpensesInput, ctx *gin.Context) *accounting.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{accounting.AuthorizationFailed},
		}
	}

	expenses, err := s.repo.Create(input, claims.Subject)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	return &accounting.BaseResult{
		Status: http.StatusOK,
		Result: expenses,
	}
}

// GetAll implements services.ExpensesService.
func (s *service) GetAll(ctx *gin.Context) *accounting.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{accounting.AuthorizationFailed},
		}
	}

	exp, err := s.repo.GetAll(claims.Subject)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	return &accounting.BaseResult{
		Status:      http.StatusOK,
		ResultCount: int64(len(exp)),
		Result:      exp,
	}
}
