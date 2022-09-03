package main

import (
	"log"
)

func main() {
	usecase, err := NewCreateOrderUsecaseProvider(container.injector)
	if err != nil {
		log.Fatal(err)
	}
	usecase.Run()

	usecase, err = NewCreateOrderUsecaseProvider(container.injector)
	if err != nil {
		log.Fatal(err)
	}
	usecase.Run()
}
