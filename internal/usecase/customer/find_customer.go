package customer

import (
	"context"
	"errors"
	"github.com/hashicorp/go-multierror"
	"golang-demo/domain"
	"golang-demo/pkg/exceptions"
)

func (uc *CustomerUseCase) GetCustomer(ctx context.Context) (*domain.Customer, *exceptions.CustomerError) {
	var multiErr *multierror.Error

	customer, err := uc.repoCustomer.GetCustomer(ctx)
	if err != nil {
		multiErr = multierror.Append(multiErr, errors.New("data not found"))
		return nil, &exceptions.CustomerError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multiErr,
		}
	}

	return customer, nil
}
