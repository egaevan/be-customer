package customer_psql

import (
	"context"
	"golang-demo/domain"
	"golang-demo/internal/repository/mapper"
	"golang-demo/internal/repository/models"
)

func (r *psqlCustomerRepository) GetCustomer(ctx context.Context) (*domain.Customer, error) {
	var data *models.Customer
	result := r.db.Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return mapper.ToDomainCustomer(data), nil
}
