package account

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/metrics"
	"github.com/ppeymann/accounting.git"
	"github.com/ppeymann/accounting.git/services"
)

type instrumentingServices struct {
	requestCounter metrics.Counter
	requestLatency metrics.Histogram
	next           services.AccountService
}

func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, services services.AccountService) services.AccountService {
	return &instrumentingServices{
		requestCounter: counter,
		requestLatency: latency,
		next:           services,
	}
}

// SignUp implements services.AccountServices.
func (i *instrumentingServices) SignUp(input *services.LoginInputDTO, ctx *gin.Context) *accounting.BaseResult {
	defer func(begin time.Time) {
		i.requestCounter.With("method", "SignUp").Add(1)
		i.requestLatency.With("method", "SignUp").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.SignUp(input, ctx)
}

// SignIn implements services.AccountService.
func (i *instrumentingServices) SignIn(input *services.LoginInputDTO, ctx *gin.Context) *accounting.BaseResult {
	defer func(begin time.Time) {
		i.requestCounter.With("method", "SignIn").Add(1)
		i.requestLatency.With("method", "SignIn").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.SignIn(input, ctx)
}

// ChangeName implements services.AccountService.
func (i *instrumentingServices) ChangeName(input *services.NameInput, ctx *gin.Context) *accounting.BaseResult {
	defer func(begin time.Time) {
		i.requestCounter.With("method", "ChangeName").Add(1)
		i.requestLatency.With("method", "ChangeName").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.ChangeName(input, ctx)
}

// ChangeCurrency implements services.AccountService.
func (i *instrumentingServices) ChangeCurrency(input *services.ChangeCurrencyInput, ctx *gin.Context) *accounting.BaseResult {
	defer func(begin time.Time) {
		i.requestCounter.With("method", "ChangeCurrency").Add(1)
		i.requestLatency.With("method", "ChangeCurrency").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.ChangeCurrency(input, ctx)
}

// GetAccount implements services.AccountService.
func (i *instrumentingServices) GetAccount(ctx *gin.Context) *accounting.BaseResult {
	defer func(begin time.Time) {
		i.requestCounter.With("method", "GetAccount").Add(1)
		i.requestLatency.With("method", "GetAccount").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.GetAccount(ctx)
}

// ChangePassword implements services.AccountService.
func (i *instrumentingServices) ChangePassword(input *services.ChangePasswordInput, ctx *gin.Context) *accounting.BaseResult {
	defer func(begin time.Time) {
		i.requestCounter.With("method", "ChangePassword").Add(1)
		i.requestLatency.With("method", "ChangePassword").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.ChangePassword(input, ctx)
}
