package customer_psql

import "gorm.io/gorm"

type psqlCustomerRepository struct {
	db *gorm.DB
}

func NewPsqlCustomerRepository(db *gorm.DB) *psqlCustomerRepository {
	return &psqlCustomerRepository{
		db: db,
	}
}
