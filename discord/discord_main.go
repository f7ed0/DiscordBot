package discord

import (
	"discordRatio/lg"
	"os"
	"os/signal"
)

func main() {
	Main()
}

func Main() {
	bot, err := NewDiscordBot(os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		lg.Error.Fatalln(err.Error())
	}
	/* 	inf, sess, err := edt.UphfLogin("ssausse_", "niiGK9W63Vy9wgR!!")
	   	if err != nil {
	   		lg.Error.Fatalln(err)
	   	}
	   	_ = sess
	   	jar, err := cookiejar.New(nil)
	   	if err != nil {
	   		return
	   	}
	   	jar.SetCookies(inf, sess)
	   	err, edt_w := edt.FetchWeek(1, jar)
	   	if err != nil {
	   		return
	   	}
	   	lg.Debug.Println(edt_w) */
	bot.Init()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	lg.Info.Println("Press Ctrl+C to exit")
	<-stop
	bot.Stop()
}
