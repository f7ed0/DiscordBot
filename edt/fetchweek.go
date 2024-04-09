package edt

import (
	"discordRatio/lg"
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

func FetchWeek(weekdiff int, jar http.CookieJar) (err error, edt []EdtElement) {
	edt = []EdtElement{}
	c := http.Client{
		Jar: jar,
	}
	c.Jar = jar
	info, err := c.Get("https://vtmob.uphf.fr/esup-vtclient-up4/stylesheets/desktop/welcome.xhtml;jsessionid=6881661BC6DDF73F957FC072FE9959D3")
	if err != nil {
		return
	}
	lg.Debug.Println(info.Request.URL.String())
	if weekdiff > 0 {
		for i := 0; i < weekdiff-1; i++ {
			strr := fmt.Sprintf("!%v", i+1)
			_, err = c.PostForm(
				"https://vtmob.uphf.fr/esup-vtclient-up4/stylesheets/desktop/welcome.xhtml",
				url.Values{
					"org.apache.myfaces.trinidad.faces.FORM": {"form_week"},
					"_noJavaScript":                          {"false"},
					"javax.faces.ViewState":                  {strr},
					"form_week:_idcl":                        {"form_week:btn_next"},
				},
			)
			if err != nil {
				return
			}
		}
		strr := fmt.Sprintf("!%v", weekdiff)
		info, err = c.PostForm(
			"https://vtmob.uphf.fr/esup-vtclient-up4/stylesheets/desktop/welcome.xhtml",
			url.Values{
				"org.apache.myfaces.trinidad.faces.FORM": {"form_week"},
				"_noJavaScript":                          {"false"},
				"javax.faces.ViewState":                  {strr},
				"form_week:_idcl":                        {"form_week:btn_next"},
			},
		)
		if err != nil {
			return
		}
	}
	a := html.NewTokenizer(info.Body)
	edt = Tinkering(a, weekdiff)
	return
}
