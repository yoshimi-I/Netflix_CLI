package cookie

import (
	"fmt"
	"net/http"
)

func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	webPage := "https://www.netflix.com/settings/viewed/BPTKRBAK5ZC6PFLO6OTBD2EF6A"
	cookies := GetCookieHandler(webPage)
	for i, cok := range cookies {
		fmt.Println(i, cok)
		http.SetCookie(w, cok)
	}
	w.WriteHeader(http.StatusOK)
}
