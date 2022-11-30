package repo

import (
	pbc "exam/customer_service/genproto/customer"
)

type CustomerStorageI interface {
	Create(*pbc.CustomerRequest) (*pbc.Customer, error)
	GetCustomer(id int) (*pbc.Customer, error)
	DeleteCustomer(id int) error
	GetCustomerList() (*pbc.CustomerListResponse, error)
	UpdateCustomer(*pbc.Customer) (*pbc.Customer, error)
}
