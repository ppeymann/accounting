package bank

import (
	"github.com/gin-gonic/gin"
	accounting "github.com/ppeymann/accounting.git"
	"github.com/ppeymann/accounting.git/services"
	validations "github.com/ppeymann/accounting.git/validation"
)

type validationService struct {
	schemas map[string][]byte
	next    services.BankService
}

func NewValidationService(schemasPath string, service services.BankService) (services.BankService, error) {
	schemas := make(map[string][]byte)
	err := validations.LoadSchema(schemasPath, schemas)
	if err != nil {
		return nil, err
	}

	return &validationService{
		schemas: schemas,
		next:    service,
	}, nil
}

// Create implements services.BankService.
func (v *validationService) Create(input *services.BankAccountInput, ctx *gin.Context) *accounting.BaseResult {
	err := validations.Validate(input, v.schemas)
	if err != nil {
		return err
	}

	return v.next.Create(input, ctx)
}

// GetAllBank implements services.BankService.
func (v *validationService) GetAllBank(ctx *gin.Context) *accounting.BaseResult {
	return v.next.GetAllBank(ctx)
}
