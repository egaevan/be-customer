package domain

import (
	"context"
	"golang-demo/pkg/exceptions"
)

type Customer struct {
	CustomerId int
	LastName   string
	FirstName  string
	Address    string
}

type ContactUsRepository interface {
	GetCustomer(ctx context.Context) (*Customer, error)
}

type ContactUsUseCase interface {
	GetCustomer(ctx context.Context) (*Customer, *exceptions.CustomerError)
}
