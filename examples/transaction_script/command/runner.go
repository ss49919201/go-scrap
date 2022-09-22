package command

type Runner[T any] interface {
	Run() (model T, err error)
}
