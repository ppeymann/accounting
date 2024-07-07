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

		// SignIn is for sign in to application
		SignIn(input *LoginInputDTO, ctx *gin.Context) *accounting.BaseResult

		// ChangeName is for changing name that user set
		ChangeName(input *NameInput, ctx *gin.Context) *accounting.BaseResult

		// ChangeCurrency is for changing currency
		ChangeCurrency(input *ChangeCurrencyInput, ctx *gin.Context) *accounting.BaseResult

		// GetAccount is for getting account
		GetAccount(ctx *gin.Context) *accounting.BaseResult
	}

	// AccountRepository represents method signatures for account domain repository.
	// so any object that stratifying this interface can be used as account domain repository.
	AccountRepository interface {
		// Create for create a account with username and password
		Create(input *LoginInputDTO) (*AccountEntity, error)

		// Find is for finding user from DB with username
		Find(username string) (*AccountEntity, error)

		// Update is for updating account entity
		Update(account *AccountEntity) error

		// ChangeName is for changing name that user set
		ChangeName(name string, id uint) (*AccountEntity, error)

		// FindByID is for finding user from DB with id
		FindByID(id uint) (*AccountEntity, error)

		// ChangeCurrency is for changing currency
		ChangeCurrency(currency CurrencyType, id uint) (*AccountEntity, error)

		accounting.BaseRepository
	}

	// AccountHandler represents method signatures for account handlers.
	// so any object that stratifying this interface can be used as account handlers.
	AccountHandler interface {
		// SignUp is sign up in application http handler.
		SignUp(ctx *gin.Context)

		// SignIn is for sign in to application http handler.
		SignIn(ctx *gin.Context)

		// ChangeName is for changing name that user set http request.
		ChangeName(ctx *gin.Context)

		// ChangeCurrency is for changing currency http request
		ChangeCurrency(ctx *gin.Context)

		// GetAccount is for getting account http request
		GetAccount(ctx *gin.Context)
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
		UserName string `json:"user_name" gorm:"column:user_name;index;unique"`

		// Password
		Password string `json:"password" gorm:"password" mapstructure:"password"`

		// Email
		Email string `json:"email" gorm:"email" mapstructure:"email"`

		// Mobile
		Mobile string `json:"mobile" gorm:"mobile" mapstructure:"mobile"`

		// FullName
		FullName string `json:"full_name" gorm:"full_name" mapstructure:"full_name"`

		// Tokens list of current account active session
		Tokens []RefreshTokenEntity `json:"-" gorm:"foreignKey:AccountID;references:ID"`

		// CurrencyType is currency type between Rial, Dollar, Dinar, ...
		CurrencyType CurrencyType `json:"currency_type" gorm:"column:currency_type"`

		BankAccount []BankAccountEntity `json:"bank_account" gorm:"bank_number;foreignKey:AccountID;references:ID"`
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

	// NameInput is DTO for parsing name request params.
	//
	// swagger:model NameInput
	NameInput struct {
		// FullName
		FullName string `json:"full_name" gorm:"full_name" mapstructure:"full_name"`
	}

	// ChangeCurrencyInput is DTO for parsing name request params.
	//
	// swagger:model ChangeCurrencyInput
	ChangeCurrencyInput struct {
		// CurrencyType
		CurrencyType CurrencyType `json:"currency_type" gorm:"currency_type" mapstructure:"currency_type"`
	}

	// CurrencyType is currency type between Rial, Dollar, Dinar, ...
	CurrencyType string
)

const (
	Rial   CurrencyType = "IRT"
	Dollar CurrencyType = "USD"
	Euro   CurrencyType = "EUR"
	Dirham CurrencyType = "AED"
	Lier   CurrencyType = "TRY"
)

var AllCurrencyType = []CurrencyType{Rial, Dollar, Euro, Dirham, Lier}
