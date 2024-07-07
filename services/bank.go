package services

import (
	"github.com/gin-gonic/gin"
	"github.com/ppeymann/accounting.git"
	"gorm.io/gorm"
)

type (
	// BankService represents method signatures for api bank endpoint.
	// so any object that stratifying this interface can be used as bank service for api endpoint.
	BankService interface {
		// Create creates new bank account
		Create(input *BankAccountInput, ctx *gin.Context) *accounting.BaseResult
	}

	// BankRepository represents method signatures for bank domain repository.
	// so any object that stratifying this interface can be used as bank domain repository.
	BankRepository interface {
		// Create creates new bank account in database
		Create(input *BankAccountInput, accountID uint) (*BankAccountEntity, error)

		accounting.BaseRepository
	}

	// BankHandler represents method signatures for bank handlers.
	// so any object that stratifying this interface can be used as bank handlers.
	BankHandler interface {
		// Create creates new bank account http request.
		Create(ctx *gin.Context)
	}

	// BankEntity Contains bank information and entity
	//
	// swagger:model BankAccountEntity
	BankAccountEntity struct {
		gorm.Model `swaggerignore:"true"`

		// Name
		Name string `json:"name" gorm:"name" mapstructure:"name"`

		// BankNumber
		BankNumber int64 `json:"bank_number" gorm:"bank_number" mapstructure:"bank_number"`

		// AccountID
		AccountID uint `json:"account_id" gorm:"account_id" mapstructure:"account_id"`

		// BankSlug
		BankSlug string `json:"bank_slug" gorm:"bank_slug" mapstructure:"bank_slug"`
	}

	// BankEntity Contains bank information
	//
	// swagger:model BankEntity
	BankEntity struct {
		gorm.Model `swaggerignore:"true"`

		// Name
		Name string `json:"name" gorm:"name" mapstructure:"name"`

		// BankSlug
		BankSlug string `json:"bank_slug" gorm:"bank_slug" mapstructure:"bank_slug"`
	}

	// BankInput Contains bank input information
	//
	// swagger:model BankInput
	BankAccountInput struct {
		// Name
		Name string `json:"name" mapstructure:"name"`

		// BankNumber
		BankNumber int64 `json:"bank_number" mapstructure:"bank_number"`

		// BankSlug
		BankSlug string `json:"bank_slug" mapstructure:"bank_slug"`
	}
)
