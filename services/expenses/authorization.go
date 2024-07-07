package expenses

import "github.com/ppeymann/accounting.git/services"

type authorizationService struct {
	next services.ExpensesService
}

func NewAuthorizationService(service services.ExpensesService) services.ExpensesService {
	return &authorizationService{
		next: service,
	}
}
