package handler

import (
	"discordRatio/edt"
	"discordRatio/lg"
	"fmt"
	"net/http/cookiejar"

	"github.com/bwmarrin/discordgo"
)

func Edt(s *discordgo.Session, i *discordgo.InteractionCreate) {
	/*msg, err := s.ChannelMessageSend(i.ChannelID, "fetching...")
	if err != nil {
		lg.Error.Println(err.Error())
	} */
	err := s.InteractionRespond(
		i.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "fetching...",
			},
		},
	)
	if err != nil {
		lg.Error.Println(err.Error())
		return
	}

	inf, sess, err := edt.UphfLoginFromDiscord(i.Member.User.ID)
	if err == edt.NoLinkError {
		a := "Je ne peux pas recuperer votre EDT"
		_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{
				{
					Type:  discordgo.EmbedTypeArticle,
					Title: a,
				},
			},
		})
		if err != nil {
			lg.Error.Println(err.Error())
			return
		}
		return
	} else if err != nil {
		lg.Error.Println(err)
	}
	_ = sess
	jar, err := cookiejar.New(nil)
	if err != nil {
		return
	}
	jar.SetCookies(inf, sess)
	offset := 0
	if len(i.ApplicationCommandData().Options) > 0 {
		offset = int(i.ApplicationCommandData().Options[0].IntValue())
	}

	err, edt_w := edt.FetchWeek(offset, jar)
	if err != nil {
		return
	}
	lg.Verbose.Println(edt_w)
	embedz := []*discordgo.MessageEmbed{{Title: "Emploi tu temps de " + i.Member.Nick, Fields: []*discordgo.MessageEmbedField{}}}
	for i, day := range edt.Weekday {
		first := true
		for _, item := range edt_w {
			if item.DayOfTheWeek == i {
				if first {
					_, _month, _day := item.StartTimeStamp.Date()
					embedz[0].Fields = append(embedz[0].Fields, &discordgo.MessageEmbedField{Name: fmt.Sprintf("%v %v %v", day, _day, edt.Month[_month])})
					first = false

				}
				embedz[0].Fields[len(embedz[0].Fields)-1].Value += "**" + item.Timing + "**\n" + item.Course + "\n" + item.Room + "\n\n"
			}
		}
	}
	lg.Verbose.Println(embedz)
	_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &embedz,
	})
	if err != nil {
		lg.Error.Println(err.Error())
		return
	}
}
