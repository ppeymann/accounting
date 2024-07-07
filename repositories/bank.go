package repositories

import (
	"encoding/json"
	"fmt"
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
