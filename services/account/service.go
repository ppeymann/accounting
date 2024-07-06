package account

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ppeymann/accounting.git"
	"github.com/ppeymann/accounting.git/auth"
	"github.com/ppeymann/accounting.git/env"
	"github.com/ppeymann/accounting.git/services"
	"github.com/ppeymann/accounting.git/utils"
	"github.com/segmentio/ksuid"
)

type service struct {
	repo   services.AccountRepository
	config *accounting.Configuration
}

func NewService(repo services.AccountRepository, config *accounting.Configuration) services.AccountService {
	return &service{
		repo:   repo,
		config: config,
	}
}

// SignUp implements services.AccountService.
func (s *service) SignUp(input *services.LoginInputDTO, ctx *gin.Context) *accounting.BaseResult {
	// first hashing password if is production mode
	if env.IsProduction() {
		hash, err := utils.HashString(input.Password)
		if err != nil {
			return &accounting.BaseResult{
				Status: http.StatusOK,
				Errors: []string{err.Error()},
			}
		}

		input.Password = hash
	}

	// create account
	account, err := s.repo.Create(input)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	// add refresh to refresh table
	refresh := services.RefreshTokenEntity{
		TokenId:   ksuid.New().String(),
		UserAgent: ctx.Request.UserAgent(),
		IssuedAt:  time.Now().UTC().Unix(),
		ExpiredAt: time.Now().Add(time.Duration(s.config.JWT.RefreshExpire) * time.Minute).UTC().Unix(),
	}

	account.Tokens = append(account.Tokens, refresh)

	err = s.repo.Update(account)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	// create token and refresh token
	paseto, err := auth.NewPasetoMaker(s.config.JWT.Secret)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	tokenClaims := &auth.Claims{
		Subject:   account.ID,
		Issuer:    s.config.JWT.Issuer,
		Audience:  s.config.JWT.Audience,
		IssuedAt:  time.Unix(refresh.IssuedAt, 0),
		ExpiredAt: time.Now().Add(time.Duration(s.config.JWT.TokenExpire) * time.Minute).UTC(),
	}

	refreshClaims := &auth.Claims{
		Subject:   account.ID,
		ID:        refresh.TokenId,
		Issuer:    s.config.JWT.Issuer,
		Audience:  s.config.JWT.Audience,
		IssuedAt:  time.Unix(refresh.IssuedAt, 0),
		ExpiredAt: time.Unix(refresh.ExpiredAt, 0),
	}

	// create token string with paseto maker
	tokenStr, err := paseto.CreateToken(tokenClaims)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	refreshStr, err := paseto.CreateToken(refreshClaims)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	return &accounting.BaseResult{
		Status: http.StatusOK,
		Result: services.TokenBundleOutput{
			Token:   tokenStr,
			Refresh: refreshStr,
			Expire:  tokenClaims.ExpiredAt,
		},
	}
}

// SignIn implements services.AccountService.
func (s *service) SignIn(input *services.LoginInputDTO, ctx *gin.Context) *accounting.BaseResult {
	// finding account
	account, err := s.repo.Find(input.UserName)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	if env.IsProduction() {
		if ok := utils.CheckHashedString(input.Password, account.Password); !ok {
			return &accounting.BaseResult{
				Status: http.StatusOK,
				Errors: []string{services.ErrSignInFailed.Error()},
			}
		}
	} else {
		if input.Password != account.Password {
			return &accounting.BaseResult{
				Status: http.StatusOK,
				Errors: []string{services.ErrSignInFailed.Error()},
			}
		}
	}

	// add refresh to refresh table
	refresh := services.RefreshTokenEntity{
		TokenId:   ksuid.New().String(),
		UserAgent: ctx.Request.UserAgent(),
		IssuedAt:  time.Now().UTC().Unix(),
		ExpiredAt: time.Now().Add(time.Duration(s.config.JWT.RefreshExpire) * time.Minute).UTC().Unix(),
	}

	account.Tokens = append(account.Tokens, refresh)

	err = s.repo.Update(account)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	// create token and refresh token
	paseto, err := auth.NewPasetoMaker(s.config.JWT.Secret)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	tokenClaims := &auth.Claims{
		Subject:   account.ID,
		Issuer:    s.config.JWT.Issuer,
		Audience:  s.config.JWT.Audience,
		IssuedAt:  time.Unix(refresh.IssuedAt, 0),
		ExpiredAt: time.Now().Add(time.Duration(s.config.JWT.TokenExpire) * time.Minute).UTC(),
	}

	refreshClaims := &auth.Claims{
		Subject:   account.ID,
		ID:        refresh.TokenId,
		Issuer:    s.config.JWT.Issuer,
		Audience:  s.config.JWT.Audience,
		IssuedAt:  time.Unix(refresh.IssuedAt, 0),
		ExpiredAt: time.Unix(refresh.ExpiredAt, 0),
	}

	// create token string with paseto maker
	tokenStr, err := paseto.CreateToken(tokenClaims)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	refreshStr, err := paseto.CreateToken(refreshClaims)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	return &accounting.BaseResult{
		Status: http.StatusOK,
		Result: services.TokenBundleOutput{
			Token:   tokenStr,
			Refresh: refreshStr,
			Expire:  tokenClaims.ExpiredAt,
		},
	}
}

// ChangeName implements services.AccountService.
func (s *service) ChangeName(input *services.NameInput, ctx *gin.Context) *accounting.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{accounting.AuthorizationFailed},
		}
	}

	account, err := s.repo.ChangeName(input.FullName, claims.Subject)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	return &accounting.BaseResult{
		Status: http.StatusOK,
		Result: account,
	}
}
