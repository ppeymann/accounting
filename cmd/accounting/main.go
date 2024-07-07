package main

import (
	"fmt"
	"log"
	"os"
	"time"

	kitLog "github.com/go-kit/log"
	"github.com/ppeymann/accounting.git"
	"github.com/ppeymann/accounting.git/cmd/accounting/pkg"
	"github.com/ppeymann/accounting.git/server"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	now := time.Now().UTC()

	base := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC).Unix()
	start := time.Date(now.Year(), now.Month(), now.Day(), 7, 35, 0, 0, time.UTC).Unix()
	end := time.Date(now.Year(), now.Month(), now.Day(), 23, 30, 0, 0, time.UTC).Unix()

	fmt.Println("date:", base, "start:", start, "end:", end)

	// initializing configuration from environment variables
	config, err := accounting.NewConfiguration()
	if err != nil {
		log.Fatal(err)
		return
	}

	// connect to DB server
	db, err := gorm.Open(pg.Open(config.DSN), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatal(err)
		return
	}

	// configuring logger
	var logger kitLog.Logger
	logger = kitLog.NewJSONLogger(kitLog.NewSyncWriter(os.Stderr))
	logger = kitLog.With(logger, "ts", kitLog.DefaultTimestampUTC)

	// Service Logger
	sl := kitLog.With(logger, "component", "http")

	// Server instance
	svc := server.NewServer(sl, config)

	// --------- initializing service -----------

	// Account
	pkg.InitAccountService(db, logger, config, svc)

	// Expenses
	pkg.InitExpensesService(db, logger, config, svc)

	// Bank
	pkg.InitBankService(db, logger, config, svc)

	// listen and serve...
	svc.Listen()
}
