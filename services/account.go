package services

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ppeymann/accounting.git"
	"gorm.io/gorm"
)

// Errors
var ErrAccountExist = errors.New("account with specified params already exist")
var ErrSignInFailed = errors.New("account not found or password error")
var ErrPermissionDenied = errors.New("specified role is not available for user")
var ErrAccountNotExist = errors.New("specified account does not exist")

type (
	// AccountService represents method signatures for api account endpoint.
	// so any object that stratifying this interface can be used as account service for api endpoint.
	AccountService interface {
		// SignUp is for sign up in application
		SignUp(input *LoginInputDTO, ctx *gin.Context) *accounting.BaseResult
	}

	// AccountRepository represents method signatures for account domain repository.
	// so any object that stratifying this interface can be used as account domain repository.
	AccountRepository interface {
		// Create for create a account with username and password
		Create(input *LoginInputDTO) (*AccountEntity, error)

		// Update is for updating account entity
		Update(account *AccountEntity) error

		accounting.BaseRepository
	}

	// AccountHandler represents method signatures for account handlers.
	// so any object that stratifying this interface can be used as account handlers.
	AccountHandler interface {
		// SignUp is sign up in application http handler.
		SignUp(ctx *gin.Context)
	}

	// LoginInputDTO is DTO for parsing register and sign in request params.
	//
	// swagger:model LoginInputDTO
	LoginInputDTO struct {
		// Username is user name for sign up
		UserName string `json:"user_name" gorm:"user_name" mapstructure:"user_name"`

		// Password is password for sign up
		Password string `json:"password" gorm:"password" mapstructure:"password"`
	}

	// RefreshTokenEntity is entity to store accounts active session
	RefreshTokenEntity struct {
		gorm.Model
		AccountID uint
		TokenId   string `json:"token_id" gorm:"column:token_id;index"`
		UserAgent string `json:"user_agent" gorm:"column:user_agent"`
		IssuedAt  int64  `json:"issued_at" bson:"issued_at" gorm:"column:issued_at"`
		ExpiredAt int64  `json:"expired_at" bson:"expired_at" gorm:"column:expired_at"`
	}

	// AccountEntity Contains account info and is entity of user account that stored on database.
	//
	// swagger:model AccountEntity
	AccountEntity struct {
		gorm.Model `swaggerignore:"true"`

		// UserName
		UserName string `json:"user_name" gorm:"user_name" mapstructure:"user_name"`

		// Password
		Password string `json:"password" gorm:"password" mapstructure:"password"`

		// Email
		Email string `json:"email" gorm:"email" mapstructure:"email"`

		// Mobile
		Mobile string `json:"mobile" gorm:"mobile" mapstructure:"mobile"`

		FullName string `json:"full_name" gorm:"full_name" mapstructure:"full_name"`

		// Tokens list of current account active session
		Tokens []RefreshTokenEntity `json:"-" gorm:"foreignKey:AccountID;references:ID"`
	}

	// TokenBundleOutput Contains Token, Refresh Token, Date and Token Expire time for Login/Verify response DTO.
	//
	// swagger:model TokenBundleOutput
	TokenBundleOutput struct {
		// Token is JWT/PASETO token staring for storing in client side as access token
		Token string `json:"token"`

		// Refresh token string used for refreshing authentication and give fresh token
		Refresh string `json:"refresh"`

		// Expire time of Token and CentrifugeToken
		Expire time.Time `json:"expire"`
	}
)
