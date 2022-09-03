package main

import "github.com/samber/do"

func init() {
	container = &Container{
		injector: do.New(),
	}

	// usecase
	do.Provide(container.injector, NewCreateOrderUsecaseProvider)

	// repository
	do.ProvideValue[UserRepository](container.injector, NewUserRepository())
	do.ProvideValue[OrderRepository](container.injector, NewOrderRepository())
}

var container *Container

type Container struct {
	injector *do.Injector
}

func NewCreateOrderUsecaseProvider(i *do.Injector) (CreateOrderUsecase, error) {
	return NewCreateOrderUsecase(
		do.MustInvoke[UserRepository](i),
		do.MustInvoke[OrderRepository](i),
	), nil
}
