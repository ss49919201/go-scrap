package main

import "testing"

func Test_fibonach(t *testing.T) {
	type args struct {
		n uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{"1", args{1}, 1},
		{"2", args{2}, 1},
		{"3", args{3}, 2},
		{"4", args{4}, 3},
		{"5", args{5}, 5},
		{"6", args{6}, 8},
		{"10", args{10}, 55},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonach(tt.args.n); got != tt.want {
				t.Errorf("fibonach() = %v, want %v", got, tt.want)
			}
		})
	}
}
