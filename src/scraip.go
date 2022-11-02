package src

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/http/cookiejar"
)

//配列に入れる型を指定
type VideoInfo struct {
	Id   int
	Name string
	Date string
	Url  string
}

func Scraip(webPage string) (VideoList map[string]VideoInfo) {

	//Cookie周りの処理
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Printf("クッキーの取得に失敗しました: %s", err)
	}

	VideoList = make(map[string]VideoInfo)

	res, err := http.Get(webPage)

	if err != nil {
		log.Printf("サイトへの接続に失敗しました: %s", err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Printf("サイトのhtmlの取得に失敗しました: %s", err)
	}

	// Bodyを読み込む
	doc.Find(".profile-hub-heade").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})

	return

}
