package main

import "fmt"

type Elm interface {
	string | int
}

func Contains[T Elm](slice []T, elm T) bool {
	for _, v := range slice {
		if v == elm {
			return true
		}
	}
	return false
}

func ContainsStr(slice []string, elm string) bool {
	for _, v := range slice {
		if v == elm {
			return true
		}
	}
	return false
}

func ContainsInt(slice []int, elm int) bool {
	for _, v := range slice {
		if v == elm {
			return true
		}
	}
	return false
}

// ジェネリクスでも実現できそう
type Animal interface {
	Dog | Cat
}

type Dog struct{}

type Cat struct{}

type AnimalCollection[T Animal] struct {
	Animals []T
}

func NewAnimalCollection[T Animal](a T) AnimalCollection[T] {
	return AnimalCollection[T]{
		Animals: []T{},
	}
}

func (a *AnimalCollection[T]) Append(a2 T) {
	a.Animals = append(a.Animals, a2)
}

func main() {
	strSlice := []string{"1", "2"}
	strElm := "1"
	intSlice := []int{1, 2}
	intElm := 1

	// non generic
	fmt.Println(ContainsInt(intSlice, intElm))
	fmt.Println(ContainsStr(strSlice, strElm))

	// generic
	fmt.Println(Contains(intSlice, intElm))
	fmt.Println(Contains(strSlice, strElm))

	var a = AnimalCollection[Cat]{
		Animals: []Cat{{}},
	}

	// var a = AnimalCollection{
	// 	Animals: []Cat{{}},
	// } コンパイルエラー

	var _ = NewAnimalCollection(Cat{})

	a.Append(Cat{})
	// Append(a, Dog{}) コンパイルエラー
}
