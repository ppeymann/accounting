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
