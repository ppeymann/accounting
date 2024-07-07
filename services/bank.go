package services

import (
	"github.com/ppeymann/accounting.git"
	"gorm.io/gorm"
)

type (
	// BankService represents method signatures for api bank endpoint.
	// so any object that stratifying this interface can be used as bank service for api endpoint.
	BankService interface{}

	// BankRepository represents method signatures for bank domain repository.
	// so any object that stratifying this interface can be used as bank domain repository.
	BankRepository interface {
		accounting.BaseRepository
	}

	// BankHandler represents method signatures for bank handlers.
	// so any object that stratifying this interface can be used as bank handlers.
	BankHandler interface{}

	// BankEntity Contains bank information and entity
	BankAccountEntity struct {
		gorm.Model `swaggerignore:"true"`

		// Name
		Name string `json:"name" gorm:"name" mapstructure:"name"`

		// BankNumber
		BankNumber string `json:"bank_number" gorm:"bank_number" mapstructure:"bank_number"`

		// AccountID
		AccountID string `json:"account_id" gorm:"account_id" mapstructure:"account_id"`

		// BankSlug
		BankSlug string `json:"bank_slug" gorm:"bank_slug" mapstructure:"bank_slug"`
	}

	BankEntity struct {
		gorm.Model `swaggerignore:"true"`

		// Name
		Name string `json:"name" gorm:"name" mapstructure:"name"`

		// BankSlug
		BankSlug string `json:"bank_slug" gorm:"bank_slug" mapstructure:"bank_slug"`
	}
)
