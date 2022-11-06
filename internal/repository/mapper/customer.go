package mapper

import (
	"golang-demo/domain"
	"golang-demo/internal/repository/models"
)

func ToDomainCustomer(p *models.Customer) *domain.Customer {
	return &domain.Customer{
		CustomerId: p.Id,
		LastName:   p.LastName,
		FirstName:  p.FirstName,
		Address:    p.Address,
	}
}
