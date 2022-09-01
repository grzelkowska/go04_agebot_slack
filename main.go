package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	// "github.com/shomali11/slacker"
)

func main() {
	// godotenv
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Could not load .env file")
	}
	SLACK_BOT_TOKEN := os.Getenv("SLACK_BOT_TOKEN")
	SLACK_APP_TOKEN := os.Getenv("SLACK_APP_TOKEN")

	os.Setenv("SLACK_BOT_TOKEN", SLACK_BOT_TOKEN)
	os.Setenv("SLACK_APP_TOKEN", SLACK_APP_TOKEN)
}
