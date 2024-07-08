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

// GetAllBank implements services.BankService.
func (s *service) GetAllBank(ctx *gin.Context) *accounting.BaseResult {
	bank, err := s.repo.GetBanks()
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	return &accounting.BaseResult{
		Status:      http.StatusOK,
		Result:      bank,
		ResultCount: int64(len(bank)),
	}
}

// GetByID implements services.BankService.
func (s *service) GetByID(id uint, ctx *gin.Context) *accounting.BaseResult {
	bank, err := s.repo.GetByID(id)
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

// DeleteBankAccount implements services.BankService.
func (s *service) DeleteBankAccount(id uint, ctx *gin.Context) *accounting.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{accounting.AuthorizationFailed},
		}
	}

	err = s.repo.DeleteBankAccount(id, claims.Subject)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	return &accounting.BaseResult{
		Status: http.StatusOK,
		Result: id,
	}
}

// UpdateBankAccount implements services.BankService.
func (s *service) UpdateBankAccount(id uint, input *services.BankAccountInput, ctx *gin.Context) *accounting.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{accounting.AuthorizationFailed},
		}
	}

	bank, err := s.repo.UpdateBankAccount(id, claims.Subject, input)
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
