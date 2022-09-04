package test

import (
	"reflect"
	"sort"
	"testing"
)

var m = map[string]string{
	"りんご":    "果物",
	"みかん":    "果物",
	"パイナップル": "果物",
	"いちご":    "果物",
	"ほうれん草":  "野菜",
	"玉ねぎ":    "野菜",
	"ごぼう":    "野菜",
	"にんじん":   "野菜",
}

func findFoods(kind string) []string {
	var foods []string
	for k, v := range m {
		if v == kind {
			foods = append(foods, k)
		}
	}
	return foods
}

func Test_findFoods(t *testing.T) {
	type args struct {
		kind string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"success/fruit",
			args{"果物"},
			[]string{
				"りんご",
				"みかん",
				"パイナップル",
				"いちご",
			},
		},
		{
			"success/fruit",
			args{"果物"},
			[]string{
				"りんご",
				"みかん",
				"パイナップル",
				"いちご",
			},
		}, {
			"success/fruit",
			args{"果物"},
			[]string{
				"りんご",
				"みかん",
				"パイナップル",
				"いちご",
			},
		}, {
			"success/fruit",
			args{"果物"},
			[]string{
				"りんご",
				"みかん",
				"パイナップル",
				"いちご",
			},
		}, {
			"success/fruit",
			args{"果物"},
			[]string{
				"りんご",
				"みかん",
				"パイナップル",
				"いちご",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findFoods(tt.args.kind)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findFoods() = %v, want %v", got, tt.want)
			}
		})
	}
}
