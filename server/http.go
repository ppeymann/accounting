package server

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
)

func (s *Server) cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PATCH", "OPTION", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authenticate", "Authorization"},
		ExposeHeaders:    []string{"Origin"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

// secure sets http security options for gin frameworks
func (s *Server) secure() gin.HandlerFunc {
	return secure.New(secure.Config{
		AllowedHosts:          s.Config.Listener.AllowedHosts,
		SSLRedirect:           true,
		SSLHost:               s.Config.Listener.SSLHost,
		STSSeconds:            315360000,
		STSIncludeSubdomains:  true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'",
		IENoOpen:              true,
		ReferrerPolicy:        "strict-origin-when-corss-origin",
		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
	})
}
