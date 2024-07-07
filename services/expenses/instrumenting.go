package expenses

import (
	"github.com/go-kit/kit/metrics"
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
