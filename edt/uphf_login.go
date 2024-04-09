package edt

import (
	"discordRatio/discord/creds"
	"discordRatio/lg"
	"errors"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"

	"golang.org/x/net/html"
)

func UphfLogin(username, password string) (url *url.URL, session []*http.Cookie, err error) {
	session = nil
	url = nil
	jar, err := cookiejar.New(nil)
	if err != nil {
		return
	}
	cli := &http.Client{
		Jar: jar,
	}
	resp, err := cli.Get("https://vtmob.uphf.fr/esup-vtclient-up4/stylesheets/desktop/welcome.xhtml")
	lg.Debug.Println(cli.Jar.Cookies(resp.Request.URL))
	tokenizer := html.NewTokenizer(resp.Body)
	execution := fetchExecution(tokenizer)
	lg.Verbose.Println(execution)

	next, err := resp.Location()
	if err != nil {
		lg.Debug.Println("NO LOCATION")
		next = resp.Request.URL
	}

	lg.Debug.Println(next.String())
	resp, err = cli.PostForm(
		next.String(),
		map[string][]string{
			"username":    {username},
			"password":    {password},
			"execution":   {execution},
			"_eventId":    {"submit"},
			"geolocation": {""},
		},
	)

	lg.Debug.Println(resp.Request.Header)
	lg.Debug.Println(resp.StatusCode)
	if resp.StatusCode >= 400 {
		err = errors.New("WRONG CREDS")
		return
	}
	lg.Debug.Println(cli.Jar.Cookies(next))
	url = next
	session = cli.Jar.Cookies(next)
	f, err := os.Create("out.html")
	b := make([]byte, 100000)
	_, err = resp.Body.Read(b)
	_, err = f.Write(b)
	return
}

func fetchExecution(tokenizer *html.Tokenizer) (result string) {
	for tokenizer.Next() != html.ErrorToken {
		tok := tokenizer.Token()
		if tok.Data == "input" {
			ok := false
			result = ""
			for _, attr := range tok.Attr {
				if attr.Key == "name" && attr.Val == "execution" {
					if ok {
						return
					}
					ok = true
				}
				if attr.Key == "value" {
					result = attr.Val
					if ok {
						return
					}
					ok = true
				}
			}
		}
	}
	return
}

var NoLinkError = errors.New("NoLinkError")

func UphfLoginFromDiscord(UserID string) (url *url.URL, session []*http.Cookie, err error) {
	cred, ok := creds.Credentials[UserID]
	if !ok {
		err = NoLinkError
		return
	}
	return UphfLogin(cred.Username, cred.Password)
}

// https://vtmob.uphf.fr/esup-vtclient-up4/stylesheets/desktop/welcome.xhtml
