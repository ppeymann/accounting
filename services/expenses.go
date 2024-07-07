package services

import (
	"github.com/ppeymann/accounting.git"
	"gorm.io/gorm"
)

type (
	// ExpensesService represents method signatures for api expenses endpoint.
	// so any object that stratifying this interface can be used as expenses service for api endpoint.
	ExpensesService interface{}

	// ExpensesRepository represents method signatures for expenses domain repository.
	// so any object that stratifying this interface can be used as expenses domain repository.
	ExpensesRepository interface {
		accounting.BaseRepository
	}

	// ExpensesHandler represents method signatures for expenses handlers.
	// so any object that stratifying this interface can be used as expenses handlers.
	ExpensesHandler interface{}

	// ExpensesEntity Contains expenses information and entity
	ExpensesEntity struct {
		gorm.Model `swaggerignore:"true"`

		// Amount
		Amount float64 `json:"amount" gorm:"amount" mapstructure:"amount"`

		// DateAndTime
		DateAndTime DateAndTime `json:"date_and_time" gorm:"date_and_time" mapstructure:"date_and_time"`

		// Category
		Category string `json:"category" gorm:"category" mapstructure:"category"`

		// BankNumber
		BankNumber string `json:"bank_number" gorm:"bank_number" mapstructure:"bank_number"`

		// BankName
		BankName string `json:"bank_name" gorm:"bank_name" mapstructure:"bank_name"`

		// Notes
		Notes string `json:"notes" gorm:"notes" mapstructure:"notes"`
	}

	DateAndTime struct {
		Year   int `json:"year" gorm:"year" mapstructure:"year"`
		Month  int `json:"month" gorm:"month" mapstructure:"month"`
		Day    int `json:"day" gorm:"day" mapstructure:"day"`
		Hour   int `json:"hour" gorm:"hour" mapstructure:"hour"`
		Minute int `json:"minute" gorm:"minute" mapstructure:"minute"`
	}
)