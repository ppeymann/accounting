package services

import (
	"github.com/gin-gonic/gin"
	"github.com/ppeymann/accounting.git"
	"gorm.io/gorm"
)

type (
	// AccountService represents method signatures for api account endpoint.
	// so any object that stratifying this interface can be used as account service for api endpoint.
	AccountService interface {
		SignUp(input *LoginInputDTO, ctx *gin.Context) accounting.BaseResult
	}

	// AccountRepository represents method signatures for account domain repository.
	// so any object that stratifying this interface can be used as account domain repository.
	AccountRepository interface {
		SignUp(input *LoginInputDTO) (*AccountEntity, error)

		accounting.BaseRepository
	}

	// AccountHandler represents method signatures for account handlers.
	// so any object that stratifying this interface can be used as account handlers.
	AccountHandler interface {
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
	}
)
