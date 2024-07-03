package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/metrics"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/mssola/user_agent"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type serviceInstrumenting struct {
	os      metrics.Counter
	browser metrics.Counter
}

// newServiceInstrumenting returns a configured instance of serviceInstrumenting.
func newServiceInstrumenting() serviceInstrumenting {
	return serviceInstrumenting{
		os: kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "global",
			Name:      "os_count",
			Help:      "num of request that made by any of OS.",
		}, []string{"os"}),
		browser: kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "global",
			Name:      "browser_count",
			Help:      "num of request that made by any of Browsers.",
		}, []string{"browser"}),
	}
}

// prometheus sets prometheus handler to http server listener
func (s *Server) prometheus() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// metrics is http middleware for serviceInstrumenting global metrics
func (s *Server) metrics() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ua := user_agent.New(ctx.GetHeader("User-Agent"))
		defer func(agent *user_agent.UserAgent) {
			b, v := ua.Browser()
			s.instrumenting.os.With("os", ua.OS()).Add(1)
			s.instrumenting.browser.With("browser", fmt.Sprintf("%s, version: %s", b, v)).Add(1)
		}(ua)
		ctx.Next()
	}
}
