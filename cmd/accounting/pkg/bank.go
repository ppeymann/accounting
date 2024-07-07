package pkg

import (
	"log"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	kitLog "github.com/go-kit/log"
	"github.com/ppeymann/accounting.git"
	"github.com/ppeymann/accounting.git/repositories"
	"github.com/ppeymann/accounting.git/server"
	"github.com/ppeymann/accounting.git/services"
	"github.com/ppeymann/accounting.git/services/bank"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"gorm.io/gorm"
)

func InitBankService(db *gorm.DB, logger kitLog.Logger, config *accounting.Configuration, server *server.Server) services.BankService {
	bankRepo := repositories.NewBankRepository(db, config.Database)

	err := bankRepo.Migrate()
	if err != nil {
		log.Fatal(err)
	}

	err = bankRepo.Import("./source/bank_import.json")
	if err != nil {
		log.Println("Err for import bank data: ", err)
	}

	// bankService create service
	bankService := bank.NewService(bankRepo)

	// bank validation
	path := getSchemaPath("bank")
	bankService, err = bank.NewValidationService(path, bankService)
	if err != nil {
		log.Fatal(err)
	}

	// @Inject logging service to chain
	bankService = bank.NewLoggingServices(logger, bankService)

	// @Inject instrumenting service to chain
	bankService = bank.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "bank",
			Name:      "request_count",
			Help:      "num of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "bank",
			Name:      "request_latency_microseconds",
			Help:      "total duration of requests in microseconds.",
		}, fieldKeys),
		bankService,
	)

	// @Inject authorization service to chain
	bankService = bank.NewAuthorizationService(bankService)

	// @Inject handlers service to chain
	_ = bank.NewHandler(bankService, server)

	return bankService
}
