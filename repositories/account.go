package repositories

import (
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/ppeymann/accounting.git/services"
	"gorm.io/gorm"
)

type repository struct {
	pg       *gorm.DB
	database string
	table    string
}

func NewAccountRepository(db *gorm.DB, database string) services.AccountRepository {
	return &repository{
		pg:       db,
		database: database,
		table:    "account_entities",
	}
}

// Create implements services.AccountRepository.
func (r *repository) Create(input *services.LoginInputDTO) (*services.AccountEntity, error) {
	// make account information with Account Entity
	account := &services.AccountEntity{
		Model:        gorm.Model{},
		UserName:     input.UserName,
		Password:     input.Password,
		CurrencyType: services.Rial,
	}

	// add user name for email OR mobile
	if strings.Contains(input.UserName, "@") {
		account.Email = input.UserName
	} else if strings.Contains(input.UserName, "+98") {
		account.Mobile = input.UserName
	}

	// Create account
	err := r.pg.Transaction(func(tx *gorm.DB) error {
		if res := r.Model().Create(account).Error; res != nil {
			str := res.(*pgconn.PgError).Message
			if strings.Contains(str, "duplicate key value") {
				return services.ErrAccountExist
			}
			return res
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return account, nil
}

// Find implements services.AccountRepository.
func (r *repository) Find(username string) (*services.AccountEntity, error) {
	account := &services.AccountEntity{}

	err := r.Model().Where("user_name = ?", username).First(account).Error
	if err != nil {
		return nil, err
	}

	return account, nil
}

// Update implements services.AccountRepository.
func (r *repository) Update(account *services.AccountEntity) error {
	return r.pg.Save(&account).Error
}

// ChangeName implements services.AccountRepository.
func (r *repository) ChangeName(name string, id uint) (*services.AccountEntity, error) {

	account, err := r.FindByID(id)
	if err != nil {
		return nil, err
	}

	account.FullName = name

	err = r.Update(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

// FindByID implements services.AccountRepository.
func (r *repository) FindByID(id uint) (*services.AccountEntity, error) {
	account := &services.AccountEntity{}

	err := r.Model().Where("id = ?", id).First(account).Error
	if err != nil {
		return nil, err
	}

	return account, nil
}

// ChangeCurrency implements services.AccountRepository.
func (r *repository) ChangeCurrency(currency services.CurrencyType, id uint) (*services.AccountEntity, error) {
	account, err := r.FindByID(id)
	if err != nil {
		return nil, err
	}

	account.CurrencyType = currency

	err = r.Update(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

// Migrate implements services.AccountRepository.
func (r *repository) Migrate() error {
	err := r.pg.AutoMigrate(&services.RefreshTokenEntity{})
	if err != nil {
		return err
	}

	return r.pg.AutoMigrate(&services.AccountEntity{})
}

// Model implements services.AccountRepository.
func (r *repository) Model() *gorm.DB {
	return r.pg.Model(&services.AccountEntity{})
}

// Name implements services.AccountRepository.
func (r *repository) Name() string {
	return r.table
}
