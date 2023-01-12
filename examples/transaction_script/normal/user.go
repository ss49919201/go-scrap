package normal

import "github.com/ss49919201/go-scrap/examples/transaction_script/model"

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

func (u *UserTS) Create(name string) (*model.User, error) {
	return &model.User{
		ID:   1,
		Name: name,
	}, nil
}
