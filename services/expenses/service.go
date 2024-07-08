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

// GetPeriodTime implements services.ExpensesService.
func (s *service) GetPeriodTime(input *services.PeriodTimeInput, ctx *gin.Context) *accounting.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{accounting.AuthorizationFailed},
		}
	}

	exp, err := s.repo.GetPeriodTime(input, claims.Subject)
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

// GetInMonth implements services.ExpensesService.
func (s *service) GetInMonth(year int, month int, ctx *gin.Context) *accounting.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{accounting.AuthorizationFailed},
		}
	}

	exp, err := s.repo.GetInMonth(year, month, claims.Subject)
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

// DeleteExpenses implements services.ExpensesService.
func (s *service) DeleteExpenses(id uint, ctx *gin.Context) *accounting.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{accounting.AuthorizationFailed},
		}
	}

	deleteId, err := s.repo.DeleteExpenses(id, claims.Subject)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	return &accounting.BaseResult{
		Status: http.StatusOK,
		Result: deleteId,
	}
}

// UpdateExpenses implements services.ExpensesService.
func (s *service) UpdateExpenses(id uint, input *services.ExpensesInput, ctx *gin.Context) *accounting.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{accounting.AuthorizationFailed},
		}
	}

	updateExpenses, err := s.repo.UpdateExpenses(id, input, claims.Subject)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	return &accounting.BaseResult{
		Status: http.StatusOK,
		Result: updateExpenses,
	}
}

// GetByID implements services.ExpensesService.
func (s *service) GetByID(id uint, ctx *gin.Context) *accounting.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{accounting.AuthorizationFailed},
		}
	}

	exp, err := s.repo.GetByID(id, claims.Subject)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	return &accounting.BaseResult{
		Status: http.StatusOK,
		Result: exp,
	}
}

// GetByBankAccountID implements services.ExpensesService.
func (s *service) GetByBankAccountID(bankID uint, ctx *gin.Context) *accounting.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{accounting.AuthorizationFailed},
		}
	}

	exp, err := s.repo.GetByBankAccountID(bankID, claims.Subject)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	return &accounting.BaseResult{
		Status: http.StatusOK,
		Result: exp,
	}
}
