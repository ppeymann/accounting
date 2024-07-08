package expenses

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
	next   services.ExpensesService
}

func NewLoggingServices(logger log.Logger, services services.ExpensesService) services.ExpensesService {
	return &loggingService{
		logger: logger,
		next:   services,
	}
}

// Create implements services.ExpensesService.
func (l *loggingService) Create(input *services.ExpensesInput, ctx *gin.Context) (result *accounting.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "Create",
			"errors", strings.Join(result.Errors, " ,"),
			"result", result,
			"input", input,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.Create(input, ctx)
}

// GetAll implements services.ExpensesService.
func (l *loggingService) GetAll(ctx *gin.Context) (result *accounting.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "GetAll",
			"errors", strings.Join(result.Errors, " ,"),
			"result", result,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.GetAll(ctx)
}

// GetPeriodTime implements services.ExpensesService.
func (l *loggingService) GetPeriodTime(input *services.PeriodTimeInput, ctx *gin.Context) (result *accounting.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "GetPeriodTime",
			"errors", strings.Join(result.Errors, " ,"),
			"result", result,
			"input", input,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.GetPeriodTime(input, ctx)
}

// GetInMonth implements services.ExpensesService.
func (l *loggingService) GetInMonth(year int, month int, ctx *gin.Context) (result *accounting.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "GetInMonth",
			"errors", strings.Join(result.Errors, " ,"),
			"result", result,
			"year", year,
			"month", month,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.GetInMonth(year, month, ctx)
}

// DeleteExpenses implements services.ExpensesService.
func (l *loggingService) DeleteExpenses(id uint, ctx *gin.Context) (result *accounting.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "DeleteExpenses",
			"errors", strings.Join(result.Errors, " ,"),
			"result", result,
			"id", id,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.DeleteExpenses(id, ctx)
}

// UpdateExpenses implements services.ExpensesService.
func (l *loggingService) UpdateExpenses(id uint, input *services.ExpensesInput, ctx *gin.Context) (result *accounting.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "UpdateExpenses",
			"errors", strings.Join(result.Errors, " ,"),
			"result", result,
			"input", input,
			"id", id,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.UpdateExpenses(id, input, ctx)
}

// GetByID implements services.ExpensesService.
func (l *loggingService) GetByID(id uint, ctx *gin.Context) (result *accounting.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "GetByID",
			"errors", strings.Join(result.Errors, " ,"),
			"result", result,
			"id", id,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.GetByID(id, ctx)
}
