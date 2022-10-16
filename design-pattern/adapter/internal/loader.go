package internal

type Loader interface {
	Load()
	Store(v any)
}
