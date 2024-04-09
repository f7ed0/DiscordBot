package edt

import (
	"discordRatio/lg"
	"strings"

	"golang.org/x/net/html"
)

var days = []string{"LUNDI", "MARDI", "MERCREDI", "JEUDI", "VENDREDI", "SAMEDI", "DIMANCHE"}

type ptarr []*string

func (pts *ptarr) generatePtarr(el *EdtElement) {
	*pts = ptarr{
		&el.Course,
		&el.Timing,
		&el.Room,
		&el.Teacher,
		&el.Group,
		&el.Id,
	}
}

func Tinkering(a *html.Tokenizer, wos int) []EdtElement {
	table_level := 0
	table_count := []int{0, 0}
	info := -1
	tr_level := 0
	tr_count := 0
	div_level := 0
	old := -1
	ct := 0
	edt := []EdtElement{}
	ptrs := ptarr{}
	for a.Next() != html.ErrorToken {
		//tag = a.Token().Type

		token := a.Token()
		/* 		if strings.Contains(token.String(), "info") {
			lg.Verbose.Println(token.String(), "|")
		} */

		switch token.Type {
		case html.StartTagToken:
			switch token.Data {
			case "table":
				table_level++
			case "tr":
				tr_level++
				//lg.Debug.Println(tr_level)
			case "div":
				div_level++
				if strings.Contains(token.String(), "content_bulle") {
					info = div_level
				}
			}
		case html.EndTagToken:
			switch token.Data {
			case "table":
				table_level--
				if table_level == 0 {
					table_count[0]++
					table_count[1] = 0
					tr_count = 0
				} else if table_level == 1 {
					table_count[1]++
				}
			case "tr":
				tr_level--
				if tr_level == 0 {
					tr_count++
				}
			case "div":
				div_level--
				if div_level < info {
					info = -1
				}
			}
		default:
			//lg.Verbose.Println(tr_level, table_level, "?", token.String())
			if table_level >= 1 && tr_level >= 2 && table_count[0] >= 3 && token.Data != "br" && info >= 0 && !strings.HasSuffix(token.String(), ": ") {
				if old != table_count[1] {
					if len(edt) > 0 {
						edt[len(edt)-1].generateTimeStamp(wos)
						lg.Verbose.Println(edt[len(edt)-1].Timing, edt[len(edt)-1].StartTimeStamp, edt[len(edt)-1].EndTimeStamp)
					}
					edt = append(edt, EdtElement{})
					ptrs.generatePtarr(&edt[len(edt)-1])
					ct = 0
					old = table_count[1]
					edt[len(edt)-1].DayOfTheWeek = (tr_count - 2) / 2
				}

				//lg.Verbose.Println("!!", table_count[1], days[(tr_count-2)/2], token.String())
				*ptrs[ct] = token.String()
				ct++
			}
		}

	}
	if len(edt) > 0 {
		edt[len(edt)-1].generateTimeStamp(wos)
	}
	//lg.Debug.Println(edt)
	return edt
}
