package main

type UserRepository interface {
	ExistUser(id int) (bool, error)
}

type userRepositoryImpl struct{}

func NewUserRepository() *userRepositoryImpl {
	return &userRepositoryImpl{}
}

func (u *userRepositoryImpl) ExistUser(id int) (bool, error) {
	return true, nil
}
