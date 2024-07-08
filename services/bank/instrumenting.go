package bank

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
	next           services.BankService
}

func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, services services.BankService) services.BankService {
	return &instrumentingService{
		requestCounter: counter,
		requestLatency: latency,
		next:           services,
	}
}

// Create implements services.BankService.
func (i *instrumentingService) Create(input *services.BankAccountInput, ctx *gin.Context) *accounting.BaseResult {
	defer func(begin time.Time) {
		i.requestCounter.With("method", "Create").Add(1)
		i.requestLatency.With("method", "Create").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.Create(input, ctx)
}

// GetAllBank implements services.BankService.
func (i *instrumentingService) GetAllBank(ctx *gin.Context) *accounting.BaseResult {
	defer func(begin time.Time) {
		i.requestCounter.With("method", "GetAllBank").Add(1)
		i.requestLatency.With("method", "GetAllBank").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.GetAllBank(ctx)
}

// GetByID implements services.BankService.
func (i *instrumentingService) GetByID(id uint, ctx *gin.Context) *accounting.BaseResult {
	defer func(begin time.Time) {
		i.requestCounter.With("method", "GetByID").Add(1)
		i.requestLatency.With("method", "GetByID").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.GetByID(id, ctx)
}

// DeleteBankAccount implements services.BankService.
func (i *instrumentingService) DeleteBankAccount(id uint, ctx *gin.Context) *accounting.BaseResult {
	defer func(begin time.Time) {
		i.requestCounter.With("method", "DeleteBankAccount").Add(1)
		i.requestLatency.With("method", "DeleteBankAccount").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.DeleteBankAccount(id, ctx)
}

// UpdateBankAccount implements services.BankService.
func (i *instrumentingService) UpdateBankAccount(id uint, input *services.BankAccountInput, ctx *gin.Context) *accounting.BaseResult {
	defer func(begin time.Time) {
		i.requestCounter.With("method", "UpdateBankAccount").Add(1)
		i.requestLatency.With("method", "UpdateBankAccount").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.UpdateBankAccount(id, input, ctx)
}
