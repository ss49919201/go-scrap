package main

import "fmt"

// 複数の構造を隠蔽した窓口を用意して
// Easy に扱えるようにする

var users = map[int]user{
	1: {"Alice"},
	2: {"Bob"},
}

var posts = make(map[int]post, 0)

type user struct {
	Name string
}

type post struct {
	userID int
	text   string
}

// generatePost 指定されたユーザーIDの投稿を作るFacade
func generatePost(userID int, text string) {
	// 存在チェック
	_, ok := users[userID]
	if !ok {
		panic("user not found")
	}
	// postを保存する
	posts[userID] = post{userID, text}
}

func main() {
	generatePost(1, "Hello, World!")
	fmt.Println(posts[1])
}
