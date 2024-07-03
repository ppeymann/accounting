package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/ppeymann/accounting.git/auth"
)

const ContextUserKey = "CONTEXT_USER"
const ContextRoleKey = "CONTEXT_ROLE"
const ContextHostKey = "Host"

var ErrUserPrincipalsNotFount = errors.New("UserPrincipals not found in context")

func CatchClaims(ctx *gin.Context, claims *auth.Claims) error {
	user, ok := ctx.Get(ContextUserKey)
	if !ok {
		return errors.New("user not found in context")
	}

	// parse context object to claims
	err := mapstructure.Decode(user, claims)
	if err != nil {
		return errors.New("error parsing user object to claims")
	}

	return nil
}
