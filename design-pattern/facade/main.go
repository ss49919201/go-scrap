package main

import "fmt"

// 複数のクラスを使う処理を一つのインターフェースにまとめる

type NotificationEmail struct{}

func (n *NotificationEmail) Send(userID string) {
	new(EmailClient).
		Send(
			getUser(userID).Email,
		)
}

type email string

type User struct {
	ID    string
	Email email
}

type EmailClient struct{}

func (e *EmailClient) Send(v email) {
	fmt.Println("Send email to", v)
}

func getUser(id string) *User {
	return &User{
		ID:    id,
		Email: "test@example.com",
	}
}
