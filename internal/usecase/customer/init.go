package customer

import "golang-demo/domain"

type CustomerUseCase struct {
	repoCustomer domain.ContactUsRepository
}

func NewCustomerUseCase(
	repoCustomer domain.ContactUsRepository,
) domain.ContactUsUseCase {
	return &CustomerUseCase{
		repoCustomer: repoCustomer,
	}
}
