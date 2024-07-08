package services

import (
	"github.com/gin-gonic/gin"
	"github.com/ppeymann/accounting.git"
	"gorm.io/gorm"
)

type (
	// ExpensesService represents method signatures for api expenses endpoint.
	// so any object that stratifying this interface can be used as expenses service for api endpoint.
	ExpensesService interface {
		// Create is for create a expenses information
		Create(input *ExpensesInput, ctx *gin.Context) *accounting.BaseResult

		// GetAll is for getting all expenses
		GetAll(ctx *gin.Context) *accounting.BaseResult

		// GetPeriodTime is for get expenses in the period of time
		GetPeriodTime(input *PeriodTimeInput, ctx *gin.Context) *accounting.BaseResult

		// GetInMonth is for get expenses in the specified month
		GetInMonth(year int, month int, ctx *gin.Context) *accounting.BaseResult

		// DeleteExpenses is for deleting a expenses
		DeleteExpenses(id uint, ctx *gin.Context) *accounting.BaseResult

		// UpdateExpenses
		UpdateExpenses(id uint, input *ExpensesInput, ctx *gin.Context) *accounting.BaseResult

		// GetByID is for get expenses by id
		GetByID(id uint, ctx *gin.Context) *accounting.BaseResult
	}

	// ExpensesRepository represents method signatures for expenses domain repository.
	// so any object that stratifying this interface can be used as expenses domain repository.
	ExpensesRepository interface {
		// Create is for create a expenses information and stored in DB
		Create(input *ExpensesInput, userID uint) (*ExpensesEntity, error)

		// GetAll is for getting all expenses from db and send it to service
		GetAll(account_id uint) ([]ExpensesEntity, error)

		// GetPeriodTime is for get expenses in the period of time
		GetPeriodTime(input *PeriodTimeInput, accountID uint) ([]ExpensesEntity, error)

		// GetInMonth is for get expenses in the specified month
		GetInMonth(year int, month int, accountID uint) ([]ExpensesEntity, error)

		// DeleteExpenses is for deleting a expenses from database
		DeleteExpenses(id uint, accountID uint) (*uint, error)

		// UpdateExpenses
		UpdateExpenses(id uint, input *ExpensesInput, accountID uint) (*ExpensesEntity, error)

		// GetByID is for get expenses by id
		GetByID(id uint, accountID uint) (*ExpensesEntity, error)

		// Update is for update expenses
		Update(exp *ExpensesEntity) error

		accounting.BaseRepository
	}

	// ExpensesHandler represents method signatures for expenses handlers.
	// so any object that stratifying this interface can be used as expenses handlers.
	ExpensesHandler interface {
		// Create is for create a expenses information http request
		Create(ctx *gin.Context)

		// GetAll is for getting all expenses http request.
		GetAll(ctx *gin.Context)

		// GetPeriodTime is for get expenses in the period of time http request.
		GetPeriodTime(ctx *gin.Context)

		// GetInMonth is for get expenses in the specified month http request.
		GetInMonth(ctx *gin.Context)

		// DeleteExpenses is for deleting expenses http request.
		DeleteExpenses(ctx *gin.Context)

		// UpdateExpenses
		UpdateExpenses(ctx *gin.Context)

		// GetByID is for get expenses by id http request.
		GetByID(ctx *gin.Context)
	}

	// ExpensesEntity Contains expenses information and entity
	ExpensesEntity struct {
		gorm.Model `swaggerignore:"true"`

		// Amount
		Amount float64 `json:"amount" gorm:"amount" mapstructure:"amount"`

		// Year is year that this expenses is done
		Year int `json:"year" gorm:"year" mapstructure:"year"`

		// Month is month that this expenses is done
		Month int `json:"month" gorm:"month" mapstructure:"month"`

		// Day is day that this expenses is done
		Day int `json:"day" gorm:"day" mapstructure:"day"`

		// Hour is hour that this expenses is done
		Hour int `json:"hour" gorm:"hour" mapstructure:"hour"`

		// Minute is minute that this expenses is done
		Minute int `json:"minute" gorm:"minute" mapstructure:"minute"`

		// Category
		Category string `json:"category" gorm:"category" mapstructure:"category"`

		// BankNumber
		BankNumber int64 `json:"bank_number" gorm:"bank_number" mapstructure:"bank_number"`

		// BankName
		BankName string `json:"bank_name" gorm:"bank_name" mapstructure:"bank_name"`

		// Notes
		Notes string `json:"notes" gorm:"notes" mapstructure:"notes"`

		// AccountID
		AccountID uint `json:"account_id" gorm:"account_id" mapstructure:"account_id"`

		// BankID
		BankID uint `json:"bank_id" gorm:"bank_id" mapstructure:"bank_id"`
	}

	DateAndTime struct {
		Year   int `json:"year" gorm:"year" mapstructure:"year"`
		Month  int `json:"month" gorm:"month" mapstructure:"month"`
		Day    int `json:"day" gorm:"day" mapstructure:"day"`
		Hour   int `json:"hour" gorm:"hour" mapstructure:"hour"`
		Minute int `json:"minute" gorm:"minute" mapstructure:"minute"`
	}

	ExpensesInput struct {
		Amount     float64     `json:"amount"`
		Date       DateAndTime `json:"date"`
		Category   string      `json:"category"`
		BankNumber int64       `json:"bank_number"`
		BankName   string      `json:"bank_name"`
		BankID     uint        `json:"bank_id"`
		Note       string      `json:"note"`
	}

	PeriodTimeInput struct {
		From PeriodDateAndTime `json:"from"`
		To   PeriodDateAndTime `json:"to"`
	}

	PeriodDateAndTime struct {
		Year  int `json:"year" mapstructure:"year"`
		Month int `json:"month" mapstructure:"month"`
	}
)
