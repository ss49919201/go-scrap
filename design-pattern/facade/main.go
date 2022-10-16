package main

import "fmt"

// 複数のクラスを使う処理を一つのインターフェースにまとめる

type NotificationEmail struct{}

type User struct {
	ID    string
	Email string
}

type EmailClient struct{}

func (e *EmailClient) Send(email string) {
	fmt.Println("Send email to", email)
}

func getUser(id string) *User {
	return &User{
		ID:    id,
		Email: "test@example.com",
	}
}

func (n *NotificationEmail) Send(userID string) {
	(&EmailClient{}).Send(getUser(userID).Email)
}
