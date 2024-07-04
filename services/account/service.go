package account

import (
	"github.com/gin-gonic/gin"
	"github.com/ppeymann/accounting.git"
	"github.com/ppeymann/accounting.git/services"
)

type service struct {
	repo   services.AccountRepository
	config *accounting.Configuration
}

func NewService(repo services.AccountRepository, config *accounting.Configuration) services.AccountService {
	return &service{
		repo:   repo,
		config: config,
	}
}

// SignUp implements services.AccountService.
func (s *service) SignUp(input *services.LoginInputDTO, ctx *gin.Context) accounting.BaseResult {
	panic("unimplemented")
}
