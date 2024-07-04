package account

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/metrics"
	"github.com/ppeymann/accounting.git"
	"github.com/ppeymann/accounting.git/services"
)

type instrumentingservices struct {
	requestCounter metrics.Counter
	requestLatency metrics.Histogram
	next           services.AccountService
}

func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, services services.AccountService) services.AccountService {
	return &instrumentingservices{
		requestCounter: counter,
		requestLatency: latency,
		next:           services,
	}
}

// SignUp implements services.Accountservices.
func (i *instrumentingservices) SignUp(input *services.LoginInputDTO, ctx *gin.Context) *accounting.BaseResult {
	defer func(begin time.Time) {
		i.requestCounter.With("method", "SignUp").Add(1)
		i.requestLatency.With("method", "SignUp").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.SignUp(input, ctx)
}
