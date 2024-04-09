package handler

import (
	"discordRatio/edt"
	"discordRatio/lg"
	"net/http/cookiejar"
	"time"

	"github.com/bwmarrin/discordgo"
)

func WhatAmIDoingNow(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
	lg.Verbose.Println(edt_w)
	t := time.Now().Unix()
	for _, item := range edt_w {
		if t >= item.StartTimeStamp.Unix() && t <= item.EndTimeStamp.Unix() {
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
			return
		}
	}
	a := "T'as pas cours #chomeur"
	_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{
			{
				Type:  discordgo.EmbedTypeArticle,
				Title: a,
			},
		},
	})
}
