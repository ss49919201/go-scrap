package main

type OrderRepository interface {
	ExistOrder(id int) (bool, error)
}

type orderRepositoryImpl struct{}

func (o *orderRepositoryImpl) ExistOrder(id int) (bool, error) {
	return true, nil
}

func NewOrderRepository() *orderRepositoryImpl {
	return &orderRepositoryImpl{}
}
