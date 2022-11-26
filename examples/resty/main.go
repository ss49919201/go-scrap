package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New()
	post(*client)
	get(*client)
}

func get(client resty.Client) {
	// use raw body
	res, err := client.R().
		SetDoNotParseResponse(true).
		Get("http://localhost:12345/example")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("res.Result() == nil", res.Result() == nil)
	body := res.RawBody()
	defer body.Close()
	b, err := ioutil.ReadAll(body)
	if err == nil {
		fmt.Println("ioutil.ReadAll(body) => ", string(b))
	} else {
		fmt.Println("ioutil.ReadAll(body) => ", err)
	}
	buf := new(bytes.Buffer)
	if _, err := buf.Write(b); err != nil {
		fmt.Println("buf.Write(b) => ", err)
	}
	b2, err := ioutil.ReadAll(buf)
	if err == nil {
		fmt.Println("ioutil.ReadAll(buf)", string(b2))
	} else {
		fmt.Println("ioutil.ReadAll(buf)", err)
	}

	// use parsed response
	type tmp struct {
		String string `json:"String"`
	}
	t := &tmp{}
	res, err = client.R().
		SetResult(t).
		Get("http://localhost:12345/example")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("res.Result() == nil", res.Result() == nil)
	t2 := res.Result().(*tmp)
	fmt.Printf("res.Result().(*tmp) => %#v\n", t2)
	b, err = ioutil.ReadAll(res.RawBody())
	if err == nil {
		fmt.Println("ioutil.ReadAll(res.RawBody()) => ", string(b))
	} else {
		fmt.Println("ioutil.ReadAll(res.RawBody()) => ", err) // http: read on closed response body
	}
}

func post(client resty.Client) {
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
