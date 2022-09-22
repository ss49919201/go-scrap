package command

import "github.com/go-tips/examples/transaction_script/model"

type CreateUserTS struct {
	name string
}

func NewCreateUserTS(name string) Runner[*model.User] {
	return &CreateUserTS{name}
}

func (ts *CreateUserTS) Run() (*model.User, error) {
	return &model.User{
		ID:   1,
		Name: ts.name,
	}, nil
}
