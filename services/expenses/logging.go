package expenses

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
	accounting "github.com/ppeymann/accounting.git"
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

// Create implements services.ExpensesService.
func (l *loggingService) Create(input *services.ExpensesInput, ctx *gin.Context) (result *accounting.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "Create",
			"errors", strings.Join(result.Errors, " ,"),
			"result", result,
			"input", input,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.Create(input, ctx)
}

// GetAll implements services.ExpensesService.
func (l *loggingService) GetAll(ctx *gin.Context) (result *accounting.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "GetAll",
			"errors", strings.Join(result.Errors, " ,"),
			"result", result,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.GetAll(ctx)
}
