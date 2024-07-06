package server

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ppeymann/accounting.git/utils"
)

// Authenticate is authentication and Authenticate middleware for http request
func (s *Server) pasetoAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// catch Authenticate header from context
		ah := ctx.GetHeader("Authorization")

		// abort request if Authenticate header is empty or not provided.
		if len(ah) == 0 {
			_ = ctx.AbortWithError(http.StatusUnauthorized, errors.New("authorization header is not provided"))
			return
		}

		// Bearer token format validation
		fields := strings.Fields(ah)
		if len(fields) != 2 {
			_ = ctx.AbortWithError(http.StatusUnauthorized, errors.New("invalid Authorization header format"))
			return
		}

		at := strings.ToLower(fields[0])
		if at != "bearer" {
			_ = ctx.AbortWithError(http.StatusUnauthorized,
				fmt.Errorf("unsupported Authenticate format : %s", fields[0]))
			return
		}

		token := fields[1]
		claims, err := s.paseto.VerifyToken(token)
		if err != nil {
			_ = ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		if claims.ExpiredAt.Before(time.Now().UTC()) {
			_ = ctx.AbortWithError(http.StatusUnauthorized, errors.New("authorization token is expired"))
			return
		}

		ctx.Set(utils.ContextUserKey, claims)
		ctx.Next()
	}
}

// Authenticate is authentication and Authenticate middleware for http request
func (s *Server) Authenticate() gin.HandlerFunc {
	return s.pasetoAuth()
}
