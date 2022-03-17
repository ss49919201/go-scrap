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
}
