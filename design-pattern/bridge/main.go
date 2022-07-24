package main

type User struct{}

// このインターフェースが橋になる
type UserRepository interface {
	GetUser(id int) User
}

type UserUseCase struct {
	userRepository UserRepository
}

// コンストラクタで実装を委譲する
func NewUserUsecase(ur UserRepository) UserUseCase {
	return UserUseCase{ur}
}

type UserRepositoryImpl1 struct{}

func (u UserRepositoryImpl1) GetUser(id int) User {
	return User{}
}

type UserRepositoryImpl2 struct{}

func (u UserRepositoryImpl2) GetUser(id int) User {
	return User{}
}

func main() {}
