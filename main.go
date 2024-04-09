package main

import (
	"discordRatio/discord"
	"discordRatio/lg"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	err := lg.Init(lg.ALL, true)
	if err != nil {
		log.Panic(err)
	}
	lg.Debug.Println("Starting...")
	discord.Main()
}
