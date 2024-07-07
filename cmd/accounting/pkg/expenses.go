package pkg

import (
	"log"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	kitLog "github.com/go-kit/log"
	"github.com/ppeymann/accounting.git"
	"github.com/ppeymann/accounting.git/repositories"
	"github.com/ppeymann/accounting.git/server"
	"github.com/ppeymann/accounting.git/services"
	"github.com/ppeymann/accounting.git/services/expenses"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"gorm.io/gorm"
)

func InitExpensesService(db *gorm.DB, logger kitLog.Logger, config *accounting.Configuration, server *server.Server) services.ExpensesService {
	expensesRepo := repositories.NewExpensesRepository(db, config.Database)

	err := expensesRepo.Migrate()
	if err != nil {
		log.Fatal(err)
	}

	// expensesService create service
	expensesService := expenses.NewService(expensesRepo)

	// getting path
	path := getSchemaPath("expenses")
	expensesService, err = expenses.NewValidationService(path, expensesService)
	if err != nil {
		log.Fatal(err)
	}

	// @Inject logging service to chain
	expensesService = expenses.NewLoggingServices(kitLog.With(logger, "component", "expenses"), expensesService)
	// @Inject instrumenting service to chain
	expensesService = expenses.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "expenses",
			Name:      "request_count",
			Help:      "num of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "expenses",
			Name:      "request_latency_microseconds",
			Help:      "total duration of requests (ms).",
		}, fieldKeys),
		expensesService,
	)

	expensesService = expenses.NewAuthorizationService(expensesService)

	_ = expenses.NewHandler(expensesService, server)

	return expensesService

}
