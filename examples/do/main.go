package main

import (
	"fmt"
	"log"

	"github.com/samber/do"
)

func init() {
	container = &Container{}
	container.injector = do.New()

	do.Provide(container.injector, NewCreateOrderUsecase)
	do.ProvideNamed(container.injector, "userRepo", func(i *do.Injector) (UserRepository, error) {
		return &userRepositoryImpl{}, nil
	})
	do.ProvideNamedValue[OrderRepository](container.injector, "orderRepo", &orderRepositoryImpl{})
}

var container *Container

type Container struct {
	injector *do.Injector
}

type UserRepository interface{}

type userRepositoryImpl struct{}

type OrderRepository interface {
	existOrder(id int) (bool, error)
}

type orderRepositoryImpl struct{}

func (self *orderRepositoryImpl) existOrder(id int) (bool, error) {
	return true, nil
}

type CreateOrderUsecase interface {
	Run()
}

type createOrderUsecaseImpl struct {
	userRepo  UserRepository
	orderRepo OrderRepository
}

func (self *createOrderUsecaseImpl) Run() {
	fmt.Println("create order")
}

func NewCreateOrderUsecase(i *do.Injector) (CreateOrderUsecase, error) {
	return &createOrderUsecaseImpl{
		userRepo:  do.MustInvokeNamed[UserRepository](i, "userRepo"),
		orderRepo: do.MustInvokeNamed[OrderRepository](i, "orderRepo"),
	}, nil
}

func main() {
	usecase, err := NewCreateOrderUsecase(container.injector)
	if err != nil {
		log.Fatal(err)
	}
	usecase.Run()
}
