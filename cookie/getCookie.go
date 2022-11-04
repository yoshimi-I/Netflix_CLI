package cookie

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func GetCookieHandler(webPage string) []*http.Cookie {

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Printf("クッキーの取得に失敗しました: %s", err)
	}

	client := &http.Client{Jar: jar}
	res, err := client.Get(webPage)

	if err != nil {
		log.Printf("サイトへの接続に失敗しました: %s", err)
	}
	set_cookie_url, err := url.Parse(webPage)
	if err != nil {
		log.Printf("クッキーの取得に失敗しました: %s", err)
	}
	cookies := jar.Cookies(set_cookie_url)

	defer res.Body.Close()

	return cookies

}
