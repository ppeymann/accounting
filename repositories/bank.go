package repositories

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

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

// GetBanks implements services.BankRepository.
func (r *bankRepository) GetBanks() ([]services.BankEntity, error) {
	var banks []services.BankEntity

	err := r.pg.Model(&services.BankEntity{}).Find(&banks).Error
	if err != nil {
		return nil, err
	}

	return banks, nil
}

// GetByID implements services.BankRepository.
func (r *bankRepository) GetByID(id uint) (*services.BankAccountEntity, error) {
	bank := &services.BankAccountEntity{}

	err := r.Model().Where("id = ?", id).First(bank).Error
	if err != nil {
		return nil, err
	}

	return bank, nil
}

// DeleteBankAccount implements services.BankRepository.
func (r *bankRepository) DeleteBankAccount(id uint, accountID uint) error {
	err := r.Model().Where("id = ? AND account_id = ?", id, accountID).Delete(&services.BankAccountEntity{}).Error
	if err != nil {
		return err
	}

	return nil
}

// UpdateBankAccount implements services.BankRepository.
func (r *bankRepository) UpdateBankAccount(id uint, accountID uint, input *services.BankAccountInput) (*services.BankAccountEntity, error) {
	bank := &services.BankAccountEntity{}

	// get bank information
	err := r.Model().Where("id = ? AND account_id = ?", id, accountID).First(bank).Error
	if err != nil {
		return nil, err
	}

	// update bank information
	bank.Name = input.Name
	bank.BankNumber = input.BankNumber
	bank.BankSlug = input.BankSlug

	// Update Bank information
	err = r.Update(bank)
	if err != nil {
		return nil, err
	}

	// get all expenses that belongs to this bank
	expensesRepo := NewExpensesRepository(r.pg, r.database)
	exps, err := expensesRepo.GetByBankAccountID(bank.ID, accountID)
	if err != nil {
		log.Println(err)
	}

	// update expenses
	for _, exp := range exps {
		exp.BankName = input.Name
		exp.BankNumber = input.BankNumber

		err = expensesRepo.Update(&exp)
		if err != nil {
			log.Println(err)
		}
	}

	return bank, nil

}

// Update implements services.BankRepository.
func (r *bankRepository) Update(bank *services.BankAccountEntity) error {
	return r.pg.Save(bank).Error
}

// Import implements services.BankRepository.
func (r *bankRepository) Import(path string) error {
	// get banks
	imports, err := r.GetBanks()
	if err != nil {
		return err
	}

	// already imported
	if len(imports) > 0 {
		return nil
	}

	// open file and read it
	fp := filepath.Clean(path)
	data, err := os.ReadFile(fp)
	if err != nil {
		return err
	}

	var source []services.BankEntity

	// decode data
	err = json.Unmarshal(data, &source)
	if err != nil {
		return err
	}

	// check the data is import or Not
	if len(source) == 0 {
		return nil
	}

	for i, b := range source {
		bank := services.BankEntity{
			Model:    gorm.Model{},
			Name:     b.Name,
			BankSlug: b.BankSlug,
		}

		err = r.pg.Model(&services.BankEntity{}).Create(&bank).Error
		if err != nil {
			fmt.Println(i, " : ", b.BankSlug, "FAILED")
			return err
		}

		fmt.Println(i, " : ", b.BankSlug, "SUCCESS")
	}

	return nil
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
