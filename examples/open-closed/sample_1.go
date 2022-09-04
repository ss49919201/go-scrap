package main

import "fmt"

type user struct {
	ID    string
	Name  string
	Tasks []string
}

type emailElementA struct {
	Subject   string
	UserID    string
	UserName  string
	UserTasks []string
}

func sendEmailA() {
	user := &user{
		ID:    "1",
		Name:  "John",
		Tasks: []string{"Task1", "Task2"},
	}

	email := emailElementA{
		Subject:   "Hello",
		UserID:    user.ID,
		UserName:  user.Name,
		UserTasks: user.Tasks,
	}

	fmt.Println("sendEmail", email)
}

// メールの文言に使う要素や、題名が変わった場合、
// emailElementだけでなく、sendEmail()も修正する必要がある
// これは修正に閉じていない

// ---------------------------------------------------------

type emailElementB struct {
	Subject   string
	UserID    string
	UserName  string
	UserTasks []string
}

func buildEmailElement(u *user) emailElementB {
	return emailElementB{
		Subject:   "Hello",
		UserID:    u.ID,
		UserName:  u.Name,
		UserTasks: u.Tasks,
	}
}

func sendEmailB() {
	user := &user{
		ID:    "1",
		Name:  "John",
		Tasks: []string{"Task1", "Task2"},
	}

	email := buildEmailElement(user)

	fmt.Println("sendEmail", email)
}

// もし題名や、メールの文言に使うuserの要素が変わった場合でも修正は呼び出し側に影響しない
// これは修正に閉じている
