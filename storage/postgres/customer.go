package postgres

import (
	pbc "exam/customer_service/genproto/customer"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type customerRepo struct {
	db *sqlx.DB
}

func NewCustomerRepo(db *sqlx.DB) *customerRepo {
	return &customerRepo{db: db}
}

func (cr *customerRepo) Create(req *pbc.CustomerRequest) (*pbc.Customer, error) {
	customerResp := &pbc.Customer{}
	fmt.Println("createga kirdi")
	err := cr.db.QueryRow(`
		INSERT INTO 
		customers(first_name, last_name, bio, email, phone_number)
		values($1,$2,$3,$4,$5)
		returning id, first_name, last_name, bio, email, phone_number
		`, req.FirstName, req.LastName, req.Bio, req.Email, req.PhoneNumber).
		Scan(&customerResp.Id, &customerResp.FirstName, &customerResp.LastName, &customerResp.Bio, &customerResp.Email, &customerResp.PhoneNumber)
	fmt.Println(err)

	if err != nil {
		return &pbc.Customer{}, err
	}
	fmt.Println(err)
	// for _, address := range req.Addresses {
	// 	addressResp := &pbc.Address{}
	// 	err := cr.db.QueryRow(`INSERT INTO addresses(street, district, customer_id) values($1,$2,$3) returning id, street, district`, address.Street, address.District, customerResp.Id).
	// 		Scan(&addressResp.Id, &addressResp.Street, &addressResp.District)
	// 	if err != nil {
	// 		return &pbc.Customer{}, err
	// 	}
	// 	customerResp.Addresses = append(customerResp.Addresses, addressResp)
	// }

	return customerResp, nil
}

func (cr *customerRepo) GetCustomer(id int) (*pbc.Customer, error) {
	fmt.Println(id)
	customerResp := &pbc.Customer{}

	err := cr.db.QueryRow(`select id, first_name, last_name, bio, email, phone_number, created_at from customers where id=$1`, id).
		Scan(&customerResp.Id, &customerResp.FirstName, &customerResp.LastName, &customerResp.Bio, &customerResp.Email, &customerResp.PhoneNumber, &customerResp.CreatedAt)
	if err != nil {
		return &pbc.Customer{}, err
	}
	fmt.Println(err)

	return customerResp, nil
}

func (cr *customerRepo) DeleteCustomer(id int) error {
	_, err := cr.db.Exec(`update customers set deleted_at = $1 where id=$2`, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

func (cr *customerRepo) GetCustomerList() (*pbc.CustomerListResponse, error) {
	customers := &pbc.CustomerListResponse{}
	rows, err := cr.db.Query(`select id, first_name, last_name,bio, email, phone_number,created_at from customers where deleted_at is null`)

	if err != nil {
		return &pbc.CustomerListResponse{}, err
	}
	for rows.Next() {
		customer := &pbc.Customer{}
		err := rows.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.Bio, &customer.Email, &customer.PhoneNumber, &customer.CreatedAt)
		if err != nil {
			return &pbc.CustomerListResponse{}, err
		}
		customers.Customers = append(customers.Customers, customer)
	}
	return customers, nil
}

func (cr *customerRepo) UpdateCustomer(req *pbc.Customer) (*pbc.Customer, error) {
	customer := &pbc.Customer{}
	err := cr.db.QueryRow(`
		update customers 
		set 
		first_name=$1, last_name=$2, bio=$3, email=$4, phone_number=$5, updated_at=$6 
		where id=$7
		returning id, first_name, last_name, bio, email, phone_number, created_at, updated_at, deletet_at`,
		req.FirstName, req.LastName, req.Bio, req.Email, req.PhoneNumber, time.Now(), req.Id).
		Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.Bio, &customer.Email, &customer.PhoneNumber, &customer.CreatedAt, &customer.UpdatedAt, &customer.DeletedAt)

	if err != nil {
		return &pbc.Customer{}, err
	}
	if req.Addresses != nil {

		for _, addr := range req.Addresses {
			address := &pbc.Address{}
			fmt.Println(addr.District, addr.Street, req.Id)
			err := cr.db.QueryRow(`update addresses set street = $1,district = $2 where customer_id=$3 returning id, street, district`,
				addr.Street, addr.District, req.Id).Scan(&address.Id, &address.Street, &address.District)
			if err != nil {
				return &pbc.Customer{}, err
			}

			customer.Addresses = append(customer.Addresses, address)
		}
	}

	return customer, nil
}

func (cr *customerRepo) CheckField(req *pbc.CheckFieldReq) (*pbc.CheckFieldRes, error) {
	fmt.Println("req: ", req)
	query := fmt.Sprintf("SELECT 1 FROM customers WHERE %s=$1", req.Field)
	res := &pbc.CheckFieldRes{}
	temp := 0
	err := cr.db.QueryRow(query, req.Value).Scan(&temp)
	fmt.Println("temp: ", temp)
	if err != nil {
		res.Exists = false
		return res, nil
	}

	if temp == 1 {
		res.Exists = true
	} else {
		res.Exists = false
	}
	fmt.Println("res:", res.Exists)
	return res, nil

}
