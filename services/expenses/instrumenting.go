package expenses

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/metrics"
	accounting "github.com/ppeymann/accounting.git"
	"github.com/ppeymann/accounting.git/services"
)

type instrumentingService struct {
	requestCounter metrics.Counter
	requestLatency metrics.Histogram
	next           services.ExpensesService
}

func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, service services.ExpensesService) services.ExpensesService {
	return &instrumentingService{
		requestCounter: counter,
		requestLatency: latency,
		next:           service,
	}
}

// Create implements services.ExpensesService.
func (i *instrumentingService) Create(input *services.ExpensesInput, ctx *gin.Context) *accounting.BaseResult {
	defer func(begin time.Time) {
		i.requestCounter.With("method", "Create").Add(1)
		i.requestLatency.With("method", "Create").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.Create(input, ctx)
}

// GetAll implements services.ExpensesService.
func (i *instrumentingService) GetAll(ctx *gin.Context) *accounting.BaseResult {
	defer func(begin time.Time) {
		i.requestCounter.With("method", "GetAll").Add(1)
		i.requestLatency.With("method", "GetAll").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.GetAll(ctx)
}

// GetPeriodTime implements services.ExpensesService.
func (i *instrumentingService) GetPeriodTime(input *services.PeriodTimeInput, ctx *gin.Context) *accounting.BaseResult {
	defer func(begin time.Time) {
		i.requestCounter.With("method", "GetPeriodTime").Add(1)
		i.requestLatency.With("method", "GetPeriodTime").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.GetPeriodTime(input, ctx)
}

// GetInMonth implements services.ExpensesService.
func (i *instrumentingService) GetInMonth(year int, month int, ctx *gin.Context) *accounting.BaseResult {
	defer func(begin time.Time) {
		i.requestCounter.With("method", "GetInMonth").Add(1)
		i.requestLatency.With("method", "GetInMonth").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.GetInMonth(year, month, ctx)
}

// DeleteExpenses implements services.ExpensesService.
func (i *instrumentingService) DeleteExpenses(id uint, ctx *gin.Context) *accounting.BaseResult {
	defer func(begin time.Time) {
		i.requestCounter.With("method", "DeleteExpenses").Add(1)
		i.requestLatency.With("method", "DeleteExpenses").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.DeleteExpenses(id, ctx)
}

// UpdateExpenses implements services.ExpensesService.
func (i *instrumentingService) UpdateExpenses(id uint, input *services.ExpensesInput, ctx *gin.Context) *accounting.BaseResult {
	defer func(begin time.Time) {
		i.requestCounter.With("method", "UpdateExpenses").Add(1)
		i.requestLatency.With("method", "UpdateExpenses").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.UpdateExpenses(id, input, ctx)
}

// GetByID implements services.ExpensesService.
func (i *instrumentingService) GetByID(id uint, ctx *gin.Context) *accounting.BaseResult {
	defer func(begin time.Time) {
		i.requestCounter.With("method", "GetByID").Add(1)
		i.requestLatency.With("method", "GetByID").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.GetByID(id, ctx)
}
