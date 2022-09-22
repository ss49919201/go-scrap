package command

type TS[T any] interface {
	Run() (model T, err error)
}
