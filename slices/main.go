package main

import "golang.org/x/exp/slices"

func insert[T any](src []T, idx int, elm T) []T {
	return slices.Insert(src, idx, elm)
}
