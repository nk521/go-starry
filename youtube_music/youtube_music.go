package youtube_music

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"strings"

	"github.com/nk521/go-starry/config"
	log "github.com/nk521/go-starry/log"
)

type YoutubeMusicManager struct {
	Client *http.Client
}

func (ymm YoutubeMusicManager) _send_request(url string, data string, post bool) {
	calltype := "GET"
	if post {
		calltype = "POST"
	}
	req, err := http.NewRequest(calltype, url, strings.NewReader(data))
	if err != nil {
		log.Panicln(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:120.0) Gecko/20100101 Firefox/120.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Connection", "keep-alive")
	// req.Header.Set("Cookie", config.GetConfig().Login.Cookies)

	resp, err := ymm.Client.Do(req)
	if err != nil {
		log.Panicln(err)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("bytes.Index(body, []byte(\"Can You Hear\")): %v\n", bytes.Index(body, []byte("Can You Hear")))
	fmt.Println(resp.Cookies())
}

func Login() {
	y := YoutubeMusicManager{}
	jar, err := cookiejar.New(nil)

	// parsing cookies to be of type http.Cookie
	header := http.Header{}
	header.Add("Cookie", config.GetConfig().Login.Cookies)
	req := http.Request{Header: header}

	_url, _ := url.Parse("https://music.youtube.com")
	jar.SetCookies(_url, req.Cookies())
	if err != nil {
		log.Panicln(err)
	}
	y.Client = &http.Client{Jar: jar}
	y._send_request("https://music.youtube.com/", "", false)
}
