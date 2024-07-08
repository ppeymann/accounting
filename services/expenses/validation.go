package expenses

import (
	"github.com/gin-gonic/gin"
	accounting "github.com/ppeymann/accounting.git"
	"github.com/ppeymann/accounting.git/services"
	validations "github.com/ppeymann/accounting.git/validation"
)

type validationService struct {
	schemas map[string][]byte
	next    services.ExpensesService
}

func NewValidationService(schemaPath string, service services.ExpensesService) (services.ExpensesService, error) {
	schemas := make(map[string][]byte)
	err := validations.LoadSchema(schemaPath, schemas)
	if err != nil {
		return nil, err
	}

	return &validationService{
		schemas: schemas,
		next:    service,
	}, nil
}

// Create implements services.ExpensesService.
func (v *validationService) Create(input *services.ExpensesInput, ctx *gin.Context) *accounting.BaseResult {
	err := validations.Validate(input, v.schemas)
	if err != nil {
		return err
	}

	return v.next.Create(input, ctx)
}

// GetAll implements services.ExpensesService.
func (v *validationService) GetAll(ctx *gin.Context) *accounting.BaseResult {
	return v.next.GetAll(ctx)
}

// GetPeriodTime implements services.ExpensesService.
func (v *validationService) GetPeriodTime(input *services.PeriodTimeInput, ctx *gin.Context) *accounting.BaseResult {
	return v.next.GetPeriodTime(input, ctx)
}
