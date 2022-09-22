package command

import "github.com/go-tips/examples/transaction_script/model"

type FindUserByID struct {
	id int
}

func NewFindUserByID(id int) Runner[*model.User] {
	return &FindUserByID{id}
}

func (ts *FindUserByID) Run() (*model.User, error) {
	return &model.User{
		ID:   ts.id,
		Name: "foo",
	}, nil
}
