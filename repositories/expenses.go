package repositories

import (
	"github.com/ppeymann/accounting.git/services"
	"gorm.io/gorm"
)

type expensesRepository struct {
	pg       *gorm.DB
	database string
	table    string
}

func NewExpensesRepository(db *gorm.DB, database string) services.ExpensesRepository {
	return &expensesRepository{
		pg:       db,
		database: database,
		table:    "expenses_entities",
	}
}

// Create implements services.ExpensesRepository.
func (r *expensesRepository) Create(input *services.ExpensesInput, userID uint) (*services.ExpensesEntity, error) {
	expenses := &services.ExpensesEntity{
		Model:      gorm.Model{},
		Amount:     input.Amount,
		Year:       input.Date.Year,
		Month:      input.Date.Month,
		Day:        input.Date.Day,
		Hour:       input.Date.Hour,
		Minute:     input.Date.Minute,
		Category:   input.Category,
		BankNumber: input.BankNumber,
		BankName:   input.BankName,
		Notes:      input.Note,
	}

	if err := r.Model().Create(expenses).Error; err != nil {
		return nil, err
	}

	return expenses, nil
}

// Migrate implements services.ExpensesRepository.
func (r *expensesRepository) Migrate() error {
	return r.pg.AutoMigrate(&services.ExpensesEntity{})
}

// Model implements services.ExpensesRepository.
func (r *expensesRepository) Model() *gorm.DB {
	return r.pg.Model(&services.ExpensesEntity{})
}

// Name implements services.ExpensesRepository.
func (r *expensesRepository) Name() string {
	return r.table
}
