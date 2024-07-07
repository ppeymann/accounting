package expenses

import (
	"github.com/go-kit/log"
	"github.com/ppeymann/accounting.git/services"
)

type loggingService struct {
	logger log.Logger
	next   services.ExpensesService
}

func NewLoggingServices(logger log.Logger, services services.ExpensesService) services.ExpensesService {
	return &loggingService{
		logger: logger,
		next:   services,
	}
}
