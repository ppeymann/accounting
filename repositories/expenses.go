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
	// create expenses structure
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
		AccountID:  userID,
		BankID:     input.BankID,
		BankSlug:   input.BankSlug,
	}

	// create expenses
	if err := r.Model().Create(expenses).Error; err != nil {
		return nil, err
	}

	return expenses, nil
}

// GetAll implements services.ExpensesRepository.
func (r *expensesRepository) GetAll(accountID uint) ([]services.ExpensesEntity, error) {
	var exp []services.ExpensesEntity

	err := r.Model().Where("account_id = ?", accountID).Find(&exp).Error
	if err != nil {
		return nil, err
	}

	return exp, err
}

// GetPeriodTime implements services.ExpensesRepository.
func (r *expensesRepository) GetPeriodTime(input *services.PeriodTimeInput, accountID uint) ([]services.ExpensesEntity, error) {
	var exp []services.ExpensesEntity

	// get expenses from input that have from and to properties and each one has year and month
	err := r.Model().Where("account_id = ? AND month >= ? AND month <= ? AND year >= ? AND year <= ?", accountID, input.From.Month, input.To.Month, input.From.Year, input.To.Year).Find(&exp).Error
	if err != nil {
		return nil, err
	}

	return exp, nil
}

// GetInMonth implements services.ExpensesRepository.
func (r *expensesRepository) GetInMonth(year int, month int, accountID uint) ([]services.ExpensesEntity, error) {
	var exp []services.ExpensesEntity

	// get expenses with specified month and year
	err := r.Model().Where("account_id = ? AND month = ? AND year = ?", accountID, month, year).Find(&exp).Error
	if err != nil {
		return nil, err
	}

	return exp, nil
}

// DeleteExpenses implements services.ExpensesRepository.
func (r *expensesRepository) DeleteExpenses(id uint, accountID uint) (*uint, error) {
	// delete expenses and return id
	err := r.Model().Where("id = ? AND account_id = ?", id, accountID).Delete(&services.ExpensesEntity{}).Error
	if err != nil {
		return nil, err
	}

	return &id, nil
}

// UpdateExpenses implements services.ExpensesRepository.
func (r *expensesRepository) UpdateExpenses(id uint, input *services.ExpensesInput, accountID uint) (*services.ExpensesEntity, error) {

	// get expenses with id and account id
	exp, err := r.GetByID(id, accountID)
	if err != nil {
		return nil, err
	}

	// update expenses information
	exp.Amount = input.Amount
	exp.Category = input.Category
	exp.BankNumber = input.BankNumber
	exp.BankName = input.BankName
	exp.Notes = input.Note
	exp.Year = input.Date.Year
	exp.Month = input.Date.Month
	exp.Day = input.Date.Day
	exp.Hour = input.Date.Hour
	exp.Minute = input.Date.Minute
	exp.BankID = input.BankID

	// update expenses
	err = r.Update(exp)
	if err != nil {
		return nil, err
	}

	return exp, nil
}

// GetByID implements services.ExpensesRepository.
func (r *expensesRepository) GetByID(id uint, accountID uint) (*services.ExpensesEntity, error) {
	exp := &services.ExpensesEntity{}

	err := r.Model().Where("id = ? AND account_id = ?", id, accountID).First(exp).Error
	if err != nil {
		return nil, err
	}

	return exp, nil
}

// GetByBankAccountID implements services.ExpensesRepository.
func (r *expensesRepository) GetByBankAccountID(bankID uint, accountID uint) ([]services.ExpensesEntity, error) {
	var exp []services.ExpensesEntity

	err := r.Model().Where("bank_id = ? AND account_id = ?", bankID, accountID).Find(&exp).Error
	if err != nil {
		return nil, err
	}

	return exp, nil
}

// Update implements services.ExpensesRepository.
func (r *expensesRepository) Update(exp *services.ExpensesEntity) error {
	return r.pg.Save(exp).Error
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
