package main

import (
	"Goservices/orders"
	"context"
	"errors"
	"log"
	"time"
)

var (
	ErrInvalidParameter = errors.New("invalid parameter")
)

type mutationResolver struct {
	server *Server
}

func (r *mutationResolver) CreateAccount(ctx context.Context, input AccountInput) (*Account, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	a, err := r.server.accountClient.PostAccount(ctx, input.Name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &Account{
		ID:   a.ID,
		Name: a.Name,
	}, nil
}

func (r *mutationResolver) CreateProduct(ctx context.Context, input ProductInput) (*Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	a, err := r.server.catalogClient.PostProduct(ctx, input.Name, input.Description, input.Price)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &Product{
		ID:          a.ID,
		Name:        a.Name,
		Description: a.Description,
		Price:       a.Price,
	}, nil

}

func (r *mutationResolver) CreateOrder(ctx context.Context, input OrderInput) (*Order, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	var products []orders.OrderedProduct
	for _, p := range input.Products {
		if p.Quantity <= 0 {

			return nil, ErrInvalidParameter
		}
		products = append(products, orders.OrderedProduct{
			ID:       p.ID,
			Quantity: uint64(p.Quantity),
		})
	}
	o, err := r.server.orderClient.PostOrder(ctx, input.AccountID, products)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &Order{
		ID:         o.ID,
		CreatedAt:  o.CreatedAt,
		TotalPrice: o.TotalPrice,
	}, nil

}
