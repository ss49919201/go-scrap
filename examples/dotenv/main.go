package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	loadEnvA()

	fmt.Println(os.Getenv("hoge"))
	fmt.Println(os.Getenv("fuga"))
}

func loadEnv() {
	godotenv.Load()
}

func loadEnvA() {
	godotenv.Load("a/.env")
}
