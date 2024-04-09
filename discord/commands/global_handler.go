package commands

import (
	"discordRatio/lg"

	"github.com/bwmarrin/discordgo"
)

func GlobalCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	lg.Debug.Printf("issued %v command.\n", i.ApplicationCommandData().Name)
	if h, ok := commandsHandlers[i.ApplicationCommandData().Name]; ok {
		h(s, i)
	}
}
