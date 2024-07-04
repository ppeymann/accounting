package repositories

import (
	"strings"

	"github.com/jackc/pgconn"
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
		Model:    gorm.Model{},
		UserName: input.UserName,
		Password: input.Password,
	}

	// add user name for email OR mobile
	if strings.Contains(input.UserName, "@") {
		account.Email = input.UserName
	} else if strings.Contains(input.UserName, "+98") {
		account.Mobile = input.UserName
	}

	// Create account
	err := r.pg.Transaction(func(tx *gorm.DB) error {
		if resultErr := r.Model().Create(account).Error; resultErr != nil {
			str := resultErr.(*pgconn.PgError).Message
			if strings.Contains(str, "duplicate key value") {
				return services.ErrAccountExist
			}

			return resultErr
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return account, nil
}

// Update implements services.AccountRepository.
func (r *repository) Update(account *services.AccountEntity) error {
	return r.pg.Save(&account).Error
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
