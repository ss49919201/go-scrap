package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/goccy/go-yaml"
	"github.com/ss49919201/scv"
)

func main() {
	s := `foo: foo
bar: bar`
	var m map[string]any
	if err := yaml.Unmarshal([]byte(s), &m); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", m)

	fmt.Println(scv.ValidateMapKeys(m, scv.Snake))

	for _, v := range m {
		if reflect.ValueOf(v).Kind() == reflect.String {
			fmt.Println(scv.ValidateMapKeys(m, scv.Snake))
		}
	}
}
