package main

import "fmt"

type CreateOrderUsecase interface {
	Run()
}

type createOrderUsecaseImpl struct {
	userRepo  UserRepository
	orderRepo OrderRepository
}

func NewCreateOrderUsecase(userRepo UserRepository, orderRepo OrderRepository) CreateOrderUsecase {
	return &createOrderUsecaseImpl{
		userRepo:  userRepo,
		orderRepo: orderRepo,
	}
}

func (c *createOrderUsecaseImpl) Run() {
	fmt.Println("create order")
}
