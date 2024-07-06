package account

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
	"github.com/ppeymann/accounting.git"
	"github.com/ppeymann/accounting.git/services"
)

type loggingServices struct {
	logger log.Logger
	next   services.AccountService
}

func NewLoggingServices(logger log.Logger, services services.AccountService) services.AccountService {
	return &loggingServices{
		logger: logger,
		next:   services,
	}
}

// SignUp implements services.Accountservices.
func (l *loggingServices) SignUp(input *services.LoginInputDTO, ctx *gin.Context) (result *accounting.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "SignUp",
			"errors", strings.Join(result.Errors, " ,"),
			"input", input,
			"result", result,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.SignUp(input, ctx)
}

// SignIn implements services.AccountService.
func (l *loggingServices) SignIn(input *services.LoginInputDTO, ctx *gin.Context) (result *accounting.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "SignIn",
			"errors", strings.Join(result.Errors, " ,"),
			"input", input,
			"result", result,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.SignIn(input, ctx)
}

// ChangeName implements services.AccountService.
func (l *loggingServices) ChangeName(input *services.NameInput, ctx *gin.Context) (result *accounting.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "ChangeName",
			"errors", strings.Join(result.Errors, " ,"),
			"input", input,
			"result", result,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.ChangeName(input, ctx)
}

// ChangeCurrency implements services.AccountService.
func (l *loggingServices) ChangeCurrency(input *services.ChangeCurrencyInput, ctx *gin.Context) (result *accounting.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "ChangeCurrency",
			"errors", strings.Join(result.Errors, " ,"),
			"input", input,
			"result", result,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.ChangeCurrency(input, ctx)
}
