package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type logger struct{}

func (logger) Errorf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}
func (logger) Warnf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}
func (logger) Debugf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func main() {
	client := resty.New()
	client.SetDebug(true)
	client.SetLogger(logger{})
	_, err := client.R().
		EnableTrace().
		SetFormData(map[string]string{
			"key1": "value1",
			"key2": "value2",
		}).
		Post("http://localhost:12345")
	if err != nil {
		fmt.Println(err)
	}
}
