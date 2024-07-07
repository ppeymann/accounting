package expenses

import "github.com/ppeymann/accounting.git/services"

type service struct {
	repo services.ExpensesRepository
}

func NewService(repo services.ExpensesRepository) services.ExpensesService {
	return &service{
		repo: repo,
	}
}
