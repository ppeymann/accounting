package repositories

import (
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

// SignUp implements services.AccountRepository.
func (r *repository) SignUp(input *services.LoginInputDTO) (*services.AccountEntity, error) {
	panic("unimplemented")
}

// Migrate implements services.AccountRepository.
func (r *repository) Migrate() error {
	panic("unimplemented")
}

// Model implements services.AccountRepository.
func (r *repository) Model() *gorm.DB {
	panic("unimplemented")
}

// Name implements services.AccountRepository.
func (r *repository) Name() string {
	panic("unimplemented")
}
