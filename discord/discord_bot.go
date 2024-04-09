package discord

import (
	"discordRatio/discord/commands"
	"discordRatio/lg"

	"github.com/bwmarrin/discordgo"
)

const prefix string = "/"

type DiscordBot struct {
	session *discordgo.Session
	me      *discordgo.User
}

func NewDiscordBot(TOKEN string) (bot DiscordBot, err error) {
	bot = DiscordBot{}
	s, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		lg.Warn.Println(err.Error())
	}
	bot.session = s
	s.AddHandler(Ready)
	bot.me, err = s.User("@me")
	if err != nil {
		lg.Error.Fatalln(err.Error())
	}
	return
}

func (bot *DiscordBot) Init() {

	bot.session.Open()
	err := commands.RegisterCommands(bot.session)
	if err != nil {
		lg.Error.Fatalln(err.Error())
	}
	bot.session.AddHandler(commands.GlobalCommandHandler)
	lg.Debug.Println("INIT DONE")

}

func (bot *DiscordBot) Stop() {
	commands.DeleteCommands(bot.session)
}

func Ready(s *discordgo.Session, m *discordgo.Ready) {
	me, err := s.User("@me")
	if err != nil {
		lg.Error.Fatalln(err.Error())
	}
	lg.Info.Printf("Logged in as : %v (ID : %v)\n", me.Username, me.ID)
}
