package src

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

//配列に入れる型を指定
type VideoInfo struct {
	Id   int
	Name string
	Date string
	Url  string
}

func Scraip(webPage string) (VideoList map[string]VideoInfo) {
	webPage = "https://www.netflix.com/settings/viewed/BPTKRBAK5ZC6PFLO6OTBD2EF6A"
	u, err := url.Parse("https://www.google.com")
	if err != nil {
		fmt.Printf("urlのパースに失敗しました:%s", err)
	}
	cookies := GetCookieHandler("https://www.google.com")
	jar, err := cookiejar.New(nil)
	jar.SetCookies(u, cookies)
	if err != nil {
		log.Printf("クッキーの設定に失敗しました: %s", err)
	}
	fmt.Println(jar)
	c := &http.Client{Jar: jar}
	time.Sleep(100)
	res, _ := c.Get(webPage)
	fmt.Println(res)
	return
}
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

// メイン
