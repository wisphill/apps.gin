package main

import (
	"apps-gin/internal"
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	userID := flag.String("user", "", "User ID")
	role := flag.String("role", "user", "Role")
	flag.Parse()

	if *userID == "" {
		log.Fatal("user is required")
	}

	token, err := internal.GenerateToken(*userID, *role)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Generated Token:")
	fmt.Println(token)
}
