package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// ref: https://github.com/vikyd/go-checksum
func main() {
	file := os.Args[1]

	h := sha256.New()
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	r := ioutil.NopCloser((bytes.NewReader(data)))
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(h, r)
	r.Close()
	if err != nil {
		panic(err)
	}

	bArr := h.Sum(nil)
	hash := fmt.Sprintf("%x", bArr)

	hFinal := sha256.New()
	fmt.Fprintf(hFinal, "%x  %s\n", bArr, "go.mod")
	bArrFinal := hFinal.Sum(nil)
	hashSynthesized := fmt.Sprintf("%x", bArrFinal)

	fmt.Println(
		hash,
		base64.StdEncoding.EncodeToString(bArr),
		hashSynthesized,
		base64.StdEncoding.EncodeToString(bArrFinal),
		"h1:"+base64.StdEncoding.EncodeToString(bArrFinal),
	)
}
