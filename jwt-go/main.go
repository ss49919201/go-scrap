package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var hmacTestKey, _ = os.ReadFile("./hmacTestKey")

func main() {
	var errs error

	m := jwt.GetSigningMethod("HS256")
	_, err := m.Sign(
		"eyJ0eXAiOiJKV1QiLA0KICJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJqb2UiLA0KICJleHAiOjEzMDA4MTkzODAsDQogImh0dHA6Ly9leGFtcGxlLmNvbS9pc19yb290Ijp0cnVlfQ",
		hmacTestKey,
	)
	if err != nil {
		errs = errors.Join(errs, err)
	}

	_, err = m.Sign(
		"eyJ0eXAiOiJKV1QiLA0KICJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJqb2UiLA0KICJleHAiOjEzMDA4MTkzODAsDQogImh0dHA6Ly9leGFtcGxlLmNvbS9pc19yb290Ijp0cnVlfQ",
		"HOGE",
	)
	if err != nil {
		errs = errors.Join(errs, err)
	}

	_, err = m.Sign(
		"",
		hmacTestKey,
	)
	if err != nil {
		errs = errors.Join(errs, err)
	}

	fmt.Println(errs)
}
