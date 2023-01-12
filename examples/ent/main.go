package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ss49919201/go-scrap/examples/ent/ent"
)

func main() {
	if err := createUser(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Done!")
}

func open(fk bool) (*ent.Client, error) {
	return ent.Open("sqlite3", "./examples/ent/example.sqlite?_fk="+strconv.FormatBool(fk))
}

func createSchema() error {
	c, err := open(true)
	if err != nil {
		return err
	}
	defer c.Close()
	if err := c.Schema.Create(context.Background()); err != nil {
		return err
	}
	return nil
}

func createUser() error {
	c, err := open(false)
	if err != nil {
		return err
	}
	defer c.Close()
	user, err := c.User.Create().SetID(10).SetName("UserName").Save(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", user)
	return nil
}
