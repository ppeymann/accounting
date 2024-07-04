package account

import (
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
func (l *loggingservices) SignUp(input *services.LoginInputDTO, ctx *gin.Context) accounting.BaseResult {
	panic("unimplemented")
}
