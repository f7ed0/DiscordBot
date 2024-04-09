package edt

import (
	"crypto/sha1"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (edt EdtElement) GenerateEmbed() *discordgo.MessageEmbed {
	h := sha1.New()
	h.Write([]byte(edt.Course))
	hashValue := h.Sum(nil)
	return &discordgo.MessageEmbed{
		Type:  discordgo.EmbedTypeArticle,
		Title: edt.Course,
		Color: int(hashToInt(hashValue)),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Date",
				Value:  edt.GenerateDate(),
				Inline: true,
			},
			{
				Name:   "Horaire",
				Value:  edt.Timing,
				Inline: true,
			},
			{
				Name:   "Salle",
				Value:  edt.Room,
				Inline: false,
			},
			{
				Name:   "Prof.",
				Value:  edt.Teacher,
				Inline: true,
			},
		},
	}
}

func (edt EdtElement) GenerateDate() string {
	_, _month, _day := edt.StartTimeStamp.Date()
	return fmt.Sprintf("%v %v %v", Weekday[edt.DayOfTheWeek], _day, Month[_month])
}

func hashToInt(hash []byte) (ret int32) {
	for i := 0; i < 4; i++ {
		ret |= int32(hash[i]) >> i
	}
	return
}
