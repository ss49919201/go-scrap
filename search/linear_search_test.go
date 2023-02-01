package main

import "testing"

func TestLinearSearch(t *testing.T) {
	if a := linear([]int{1, 3, 5, 7, 9}, 3); a != 3 {
		t.Errorf("expected 3, actual %d", a)
	}
	if a := linear([]int{1, 3, 5, 7, 9}, 10); a != -1 {
		t.Errorf("expected -1, actual %d", a)
	}
}
