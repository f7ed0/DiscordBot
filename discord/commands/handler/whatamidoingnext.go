package handler

import (
	"discordRatio/edt"
	"discordRatio/lg"
	"math"
	"net/http/cookiejar"
	"time"

	"github.com/bwmarrin/discordgo"
)

func WhatAmIDoingNext(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
	lg.Debug.Println(i.Member.User)
	inf, sess, err := edt.UphfLoginFromDiscord(i.Member.User.ID)
	if err == edt.NoLinkError {
		a := "Je ne peux pas recuperer votre identifiant"
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

	jar, err := cookiejar.New(nil)
	if err != nil {
		return
	}
	jar.SetCookies(inf, sess)

	err, edt_w := edt.FetchWeek(0, jar)
	if err != nil {
		return
	}
	if len(edt_w) == 0 {
		err, edt_w = edt.FetchWeek(1, jar)
		if err != nil {
			return
		}
	}
	lg.Verbose.Println(edt_w)
	t := time.Now().Unix()
	iter := int64(math.MaxInt64)
	min := -1
	for i, item := range edt_w {
		if t-300 <= item.StartTimeStamp.Unix() && iter > item.StartTimeStamp.Unix() {
			min = i
			iter = item.StartTimeStamp.Unix()
			lg.Verbose.Println(i)
		}
	}
	if min >= 0 {
		item := edt_w[min]
		_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			//Content: &a,
			Embeds: &[]*discordgo.MessageEmbed{
				item.GenerateEmbed(),
			},
		})
		if err != nil {
			lg.Error.Println(err.Error())
			return
		}
	} else {
		a := "Je n'ai rien trouvé..."
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
	}

}
