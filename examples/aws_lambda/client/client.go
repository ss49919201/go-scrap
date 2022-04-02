package client

import (
	"fmt"
	"net/rpc"
)

type Input struct {
	A string
	b int
}

func Call() {
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	var reply struct{}
	err = client.Call("Function.Ping", &struct{}{}, &reply)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ping: %#v\n", reply)

	err = client.Call("Function.Invoke", &struct{}{}, &reply)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Invoke: %#v\n", reply)

	// panic: gob: type mismatch: no fields matched compiling decoder for InvokeRequest
	// err = client.Call("Function.Invoke", &Input{}, &reply)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Invoke: %#v\n", reply)

	client.Close()
}
