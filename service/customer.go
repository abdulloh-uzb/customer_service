package service

import (
	"context"
	pbc "exam/customer_service/genproto/customer"
	pbp "exam/customer_service/genproto/post"
	pbr "exam/customer_service/genproto/reyting"
	l "exam/customer_service/pkg/logger"
	"exam/customer_service/service/grpcClient"
	"exam/customer_service/storage"
	"fmt"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CustomerService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpcClient.GrpcClientI
}

func NewCustomerService(db *sqlx.DB, log l.Logger, client grpcClient.GrpcClientI) *CustomerService {
	return &CustomerService{
		storage: storage.NewStorage(db),
		logger:  log,
		client:  client,
	}
}

func (c *CustomerService) Create(ctx context.Context, req *pbc.CustomerRequest) (*pbc.Customer, error) {
	fmt.Println("servisdagi createga kirdi")
	customer, err := c.storage.Customer().Create(req)
	fmt.Println("servisdagi createda error bor ekan: ", err)
	if err != nil {
		c.logger.Error("error while create customer", l.Error(err))
		return &pbc.Customer{}, err
	}
	return customer, nil
}

func (c *CustomerService) GetCustomer(ctx context.Context, req *pbc.CustomerId) (*pbc.Customer, error) {
	customer, err := c.storage.Customer().GetCustomer(int(req.Id))

	if err != nil {
		c.logger.Error("error while create customer", l.Error(err))
		return &pbc.Customer{}, status.Error(codes.Internal, "something went wrong, please check get costumer")
	}

	rankings, err := c.client.Ranking().GetRankingsByCustomerId(ctx, &pbr.Id{Id: req.Id})
	if err != nil {
		return &pbc.Customer{}, err
	}
	for _, r := range rankings.Rankings {
		customer.Rankings = append(customer.Rankings, &pbc.Ranking{
			Name:        r.Name,
			Description: r.Description,
			Ranking:     r.Ranking,
			PostId:      r.PostId,
			CustomerId:  r.CustomerId,
		})
	}

	posts, err := c.client.Post().GetPostByCustomerId(ctx, &pbp.Id{Id: req.Id})
	if err != nil {
		return &pbc.Customer{}, err
	}

	for _, p := range posts.Posts {
		customer.Posts = append(customer.Posts, &pbc.Post{
			Name:        p.Name,
			Description: p.Description,
			CustomerId:  p.CustomerId,
		})
	}

	return customer, nil
}

func (c *CustomerService) DeleteCustomer(ctx context.Context, req *pbc.CustomerId) (*pbc.Empty, error) {

	err := c.storage.Customer().DeleteCustomer(int(req.Id))
	if err != nil {
		c.logger.Error("error while create customer", l.Error(err))
		return &pbc.Empty{}, err
	}
	return &pbc.Empty{}, nil
}

func (c *CustomerService) GetCustomerList(ctx context.Context, req *pbc.Empty) (*pbc.CustomerListResponse, error) {
	customers, err := c.storage.Customer().GetCustomerList()
	if err != nil {
		c.logger.Error("error while get customer", l.Error(err))
		return &pbc.CustomerListResponse{}, err
	}

	for _, cust := range customers.Customers {
		rankings, err := c.client.Ranking().GetRankingsByCustomerId(ctx, &pbr.Id{Id: cust.Id})
		if err != nil {
			return &pbc.CustomerListResponse{}, err
		}

		for _, r := range rankings.Rankings {
			cust.Rankings = append(cust.Rankings, &pbc.Ranking{
				Name:        r.Name,
				Description: r.Description,
				Ranking:     r.Ranking,
				PostId:      r.PostId,
				CustomerId:  r.CustomerId,
			})
		}

		posts, err := c.client.Post().GetPostByCustomerId(ctx, &pbp.Id{Id: cust.Id})
		if err != nil {
			return &pbc.CustomerListResponse{}, err
		}

		for _, p := range posts.Posts {
			cust.Posts = append(cust.Posts, &pbc.Post{
				Name:        p.Name,
				Description: p.Description,
				CustomerId:  p.CustomerId,
			})
		}

	}

	return customers, nil
}

func (c *CustomerService) UpdateCustomer(ctx context.Context, req *pbc.Customer) (*pbc.Customer, error) {
	customer, err := c.storage.Customer().UpdateCustomer(req)
	if err != nil {
		c.logger.Error("error while update customer", l.Error(err))
		return &pbc.Customer{}, err
	}
	return customer, nil
}
