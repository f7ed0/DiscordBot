package commands

import (
	"discordRatio/lg"

	"github.com/bwmarrin/discordgo"
)

func RegisterCommands(s *discordgo.Session) (err error) {
	RegisteredComands = make([]*discordgo.ApplicationCommand, len(commands))
	var rcmd *discordgo.ApplicationCommand
	for i, command := range commands {
		lg.Verbose.Println(command)
		rcmd, err = s.ApplicationCommandCreate(s.State.User.ID, "", command)
		if err != nil {
			lg.Error.Printf("Cannot create %v : %v\n", command.Name, err.Error())
			return
		}
		RegisteredComands[i] = rcmd
	}
	lg.Info.Printf("Registered %v commands.\n", len(RegisteredComands))
	return
}

func DeleteCommands(s *discordgo.Session) (err error) {
	lg.Debug.Println("Removing commands...")
	/* 	registeredCommands, err := s.ApplicationCommands(s.State.User.ID, "")
	   	if err != nil {
	   		log.Fatalf("Could not fetch registered commands: %v", err)
	   	} */

	for _, v := range RegisteredComands {
		err = s.ApplicationCommandDelete(s.State.User.ID, "", v.ID)
		if err != nil {
			lg.Error.Printf("Cannot delete '%v' command: %v\n", v.Name, err)
			return
		}
	}
	return
}

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "ping",
		Description: "send pong",
	},
	{
		Name:        "edt",
		Description: "Retrieve you week time table",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "weekoffset",
				Description: "the number of week separating you to the week you want to fetch",
			},
		},
	},
	{
		Name:        "whatamidoingnow",
		Description: "get your current class",
	},
	{
		Name:        "whatamidoingnext",
		Description: "get your next class",
	},
}

var RegisteredComands []*discordgo.ApplicationCommand
