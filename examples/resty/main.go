package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New()
	client.SetDebug(true)
	_, err := client.R().
		SetFormData(map[string]string{
			"key1": "value1",
			"key2": "value2",
			"キー":   "バリュー", // %E3%82%AD%E3%83%BC=%E3%83%90%E3%83%AA%E3%83%A5%E3%83%BC
		}).
		Post("http://localhost:12345/")
	if err != nil {
		fmt.Println(err)
	}
}
