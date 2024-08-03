package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	kitlog "github.com/go-kit/log"
	"github.com/ppeymann/accounting.git"
	"github.com/ppeymann/accounting.git/auth"
	"github.com/ppeymann/accounting.git/docs"
	"github.com/ppeymann/accounting.git/env"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	Router        *gin.Engine
	Config        *accounting.Configuration
	instrumenting serviceInstrumenting

	paseto auth.TokenMaker
	Logger kitlog.Logger
}

// EnvMode specified the running env 'release' represents production mode and ” represents development.
// it depended on gin GIN_MODE env for unifying and simplicity of setting.
var EnvMode = ""

func NewServer(logger kitlog.Logger, config *accounting.Configuration) *Server {
	svr := &Server{
		Logger:        logger,
		Config:        config,
		instrumenting: newServiceInstrumenting(),
	}

	router := gin.New()
	router.Use(gin.Recovery())

	// determining environment
	EnvMode = os.Getenv("GIN_MODE")

	// setting swagger info if not in production mode
	if env.GetStringDefault("SWAGGER_ENABLE", "false") == "false" {
		docs.SwaggerInfo.Title = fmt.Sprintf("Accounting Backend [ AuthMode: %s ]", "Paseto")
		docs.SwaggerInfo.Description = "The Swagger Documentation For Accounting Backend API Server"
		docs.SwaggerInfo.Version = "1.0"
		docs.SwaggerInfo.Host = env.GetStringDefault("HOST_URL", "https://accounting-be.liara.run")
		docs.SwaggerInfo.BasePath = "/api/v1"
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
	}

	// binding global metrics middleware
	router.Use(svr.metrics())

	if env.GetStringDefault("CORS_ENABLE", "false") == "false" {
		router.Use(svr.cors())
	}

	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatal(err)
	}

	svr.Router = router

	// bind prometheus to /metrics route
	svr.Router.GET("/metric", svr.prometheus())

	svr.paseto, err = auth.NewPasetoMaker(svr.Config.JWT.Secret)
	if err != nil {
		log.Fatal(err)
	}

	return svr
}

func (s *Server) Listen() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer stop()

	if env.GetStringDefault("SWAGGER_ENABLE", "false") == "false" {
		s.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	srv := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
		Addr:              fmt.Sprintf("%s:%d", s.Config.Listener.Host, s.Config.Listener.Port),
		Handler:           s.Router,
	}

	// start https server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http listener stopped : %s", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully hospital server, press Ctrl+C again to force")

	// The context is used to inform the server it has 30 seconds to finish
	// the request it is currently handling

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown", err)
	}

	log.Println("hospital service exiting")
}
