package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		n    int
		rows [][]int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"9786", args{4, [][]int64{{1, 1, 6, 1}, {1, 1, 1, 9}, {7, 1, 1, 1}, {1, 8, 1, 1}}}, 9786},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.rows); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
