package main

import "fmt"

type emailElement struct {
	Subject   string
	UserID    string
	UserName  string
	UserTasks []string
}

func sendEmail() {
	email := emailElement{
		Subject:   "Hello",
		UserID:    "1",
		UserName:  "John",
		UserTasks: []string{"Task1", "Task2"},
	}

	fmt.Println("sendEmail", email)
}

// メールの文言に使う要素が変わった場合、
// emailElementだけでなく、sendEmail()も修正する必要がある
// これは修正に閉じていない
