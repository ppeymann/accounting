package expenses

import (
	"github.com/ppeymann/accounting.git/server"
	"github.com/ppeymann/accounting.git/services"
)

type handler struct {
	service services.ExpensesService
}

func NewHandler(service services.ExpensesService, s *server.Server) services.ExpensesHandler {
	handler := &handler{
		service: service,
	}

	_ = s.Router.Group("/api/v1/expenses")
	{

	}

	return handler
}
