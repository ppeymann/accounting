package account

import (
	"github.com/gin-gonic/gin"
	"github.com/ppeymann/accounting.git/server"
	"github.com/ppeymann/accounting.git/services"
)

type handler struct {
	service services.AccountService
}

func NewHandler(svc services.AccountService, s *server.Server) services.AccountHandler {
	handler := &handler{
		service: svc,
	}

	group := s.Router.Group("/api/v1/account")
	{
		group.GET("/signup", handler.SignUp)
	}

	return handler
}

// SignUp implements services.AccountHandler.
func (h *handler) SignUp(ctx *gin.Context) {
	panic("unimplemented")
}
