package main

import (
	"discord_bot/bot"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	bot.BotToken = os.Getenv("DISCORD_KEY")
	bot.Run()
}
