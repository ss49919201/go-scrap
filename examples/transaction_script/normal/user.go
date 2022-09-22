package normal

import "github.com/go-tips/examples/transaction_script/model"

type UserTS struct{}

func NewUserTS() *UserTS {
	return &UserTS{}
}

func (u *UserTS) FindByID(id int) *model.User {
	return &model.User{
		ID:   id,
		Name: "foo",
	}
}

func (u *UserTS) Create(name string) error {
	return nil
}
