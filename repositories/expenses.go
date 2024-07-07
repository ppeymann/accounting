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
