package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mmcdole/gofeed"
)

func main() {
	parser := gofeed.NewParser()
	feed, err := parser.ParseURLWithContext("https://zenn.dev/ss49919201/feed", context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Link %s\n", feed.Link)
	fmt.Printf("Updated %s\n", feed.Updated)
}
