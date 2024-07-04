package account

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
	"github.com/ppeymann/accounting.git"
	"github.com/ppeymann/accounting.git/services"
)

type loggingservices struct {
	logger log.Logger
	next   services.AccountService
}

func NewLoggingServices(logger log.Logger, services services.AccountService) services.AccountService {
	return &loggingservices{
		logger: logger,
		next:   services,
	}
}

// SignUp implements services.Accountservices.
func (l *loggingservices) SignUp(input *services.LoginInputDTO, ctx *gin.Context) (result *accounting.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "SignUp",
			"errors", strings.Join(result.Errors, ","),
			"input", input,
			"result", result,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.SignUp(input, ctx)
}
