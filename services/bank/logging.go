package bank

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
	next   services.BankService
}

func NewLoggingServices(logger log.Logger, services services.BankService) services.BankService {
	return &loggingService{
		logger: logger,
		next:   services,
	}
}

// Create implements services.BankService.
func (l *loggingService) Create(input *services.BankAccountInput, ctx *gin.Context) (result *accounting.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "Create",
			"errors", strings.Join(result.Errors, " ,"),
			"input", input,
			"result", result,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.Create(input, ctx)
}