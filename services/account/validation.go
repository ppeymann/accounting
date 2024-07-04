package account

import (
	"github.com/gin-gonic/gin"
	accounting "github.com/ppeymann/accounting.git"
	"github.com/ppeymann/accounting.git/services"
	validations "github.com/ppeymann/accounting.git/validation"
)

type validationService struct {
	schemas map[string][]byte
	next    services.AccountService
}

func NewValidationService(schemaPath string, service services.AccountService) (services.AccountService, error) {
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

// SignUp implements services.AccountService.
func (v *validationService) SignUp(input *services.LoginInputDTO, ctx *gin.Context) accounting.BaseResult {
	panic("unimplemented")
}
