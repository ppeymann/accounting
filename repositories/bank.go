package repositories

import (
	"github.com/ppeymann/accounting.git/services"
	"gorm.io/gorm"
)

type bankRepository struct {
	pg       *gorm.DB
	database string
	table    string
}

func NewBankRepository(db *gorm.DB, database string) services.BankRepository {
	return &bankRepository{
		pg:       db,
		database: database,
		table:    "bank_account_entities",
	}
}

// Create implements services.BankRepository.
func (r *bankRepository) Create(input *services.BankAccountInput, accountID uint) (*services.BankAccountEntity, error) {

	// first check account to know is exist or NOT
	accountRepo := NewAccountRepository(r.pg, r.database)

	account, err := accountRepo.FindByID(accountID)
	if err != nil {
		return nil, err
	}

	bankAccount := &services.BankAccountEntity{
		Model:      gorm.Model{},
		Name:       input.Name,
		BankNumber: input.BankNumber,
		AccountID:  account.ID,
		BankSlug:   input.BankSlug,
	}

	err = r.pg.Transaction(func(tx *gorm.DB) error {
		if res := r.Model().Create(bankAccount).Error; err != nil {
			return res
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return bankAccount, nil
}

// Migrate implements services.BankRepository.
func (r *bankRepository) Migrate() error {
	err := r.pg.AutoMigrate(&services.BankEntity{})
	if err != nil {
		return err
	}

	return r.pg.AutoMigrate(&services.BankAccountEntity{})
}

// Model implements services.BankRepository.
func (r *bankRepository) Model() *gorm.DB {
	return r.pg.Model(&services.BankAccountEntity{})
}

// Name implements services.BankRepository.
func (r *bankRepository) Name() string {
	return r.table
}
