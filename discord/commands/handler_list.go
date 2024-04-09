package commands

import (
	"discordRatio/discord/commands/handler"
	"discordRatio/lg"

	"github.com/bwmarrin/discordgo"
)

var commandsHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"ping": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		_, err := s.ChannelMessageSend(i.ChannelID, "pong")
		if err != nil {
			lg.Warn.Println(err.Error())
		}
	},
	"edt":              handler.Edt,
	"whatamidoingnow":  handler.WhatAmIDoingNow,
	"whatamidoingnext": handler.WhatAmIDoingNext,
}
