package youtube_music

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"strings"

	"github.com/nk521/go-starry/config"
	log "github.com/nk521/go-starry/log"
)

var (
	YTM_DOMAIN     = "https://music.youtube.com"
	YTM_BASE_API   = YTM_DOMAIN + "/youtubei/v1/"
	YTM_PARAMS     = "?alt=json"
	YTM_PARAMS_KEY = "&key=AIzaSyC9XL3ZjWddXya6X74dJoCTL-WEYFDNX30"
	USER_AGENT     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:120.0) Gecko/20100101 Firefox/120.0"
)

type YoutubeMusicManager struct {
	Client    *http.Client
	CookieJar *cookiejar.Jar
	Headers   *http.Header
	Context   map[string]map[string]map[string]string
	sapisid   string
}

func (ymm YoutubeMusicManager) sendGETRequest(url string, params map[string]string) *http.Response {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Panicln(err)
	}

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	req.Header = *ymm.Headers

	resp, err := ymm.Client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	return resp
}

func (ymm YoutubeMusicManager) sendPOSTRequest(endpoint string, body map[string]string, additional_params string) []byte {
	params := YTM_PARAMS + YTM_PARAMS_KEY
	origin := ymm.Headers.Get("origin")
	if origin == "" {
		ymm.Headers.Get("x-origin")
	}
	ymm.Headers.Set("authorization", getAuthorization(ymm.sapisid+" "+origin))

	jsonStr, err := json.Marshal(body)
	if err != nil {
		log.Panicln(err)
	}
	req, err := http.NewRequest(http.MethodPost, YTM_BASE_API+endpoint+params+additional_params, strings.NewReader(string(jsonStr)))
	if err != nil {
		log.Panicln(err)
	}

	req.Header = *ymm.Headers

	resp, err := ymm.Client.Do(req)
	if err != nil {
		log.Panicln(err)
	}

	if resp.StatusCode >= 400 {
		log.Panicln(fmt.Errorf("got %s @ %s", resp.Status, req.URL.String()))
	}

	defer resp.Body.Close()
	response_body, _ := io.ReadAll(resp.Body)

	return response_body
}

func (ymm YoutubeMusicManager) Login() error {
	_, _ = ymm.Client.Get("")

	return nil
}

// func (ymm YoutubeMusicManager) prepareHeaders() error {
// 	return nil
// }

func (ymm YoutubeMusicManager) Init() {
	// set headers
	{
		ymm.Headers.Set("User-Agent", USER_AGENT)
		ymm.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
		ymm.Headers.Set("Accept-Language", "en-US,en;q=0.5")
		ymm.Headers.Set("Connection", "keep-alive")
	}

	// make cookie jar
	{
		cookies := config.GetConfig().Login.Headers
		if len(cookies) <= 0 {
			log.Panicln(fmt.Errorf("please provide cookies in `%s`", config.GetRawConfig().ConfigFileUsed()))
		}

		header := http.Header{}
		config.GetConfig().Login.Headers += ";SOCS=CAI"
		header.Add("Cookie", config.GetConfig().Login.Headers)
		req := http.Request{Header: header}
		_url, _ := url.Parse("https://music.youtube.com")

		_cookie_jar, err := cookiejar.New(nil)
		if err != nil {
			log.Panicln(err)
		}
		ymm.CookieJar = _cookie_jar
		ymm.CookieJar.SetCookies(_url, req.Cookies())
	}

	// init context
	ymm.Context = initContext()

	// and finally, the client
	ymm.Client = &http.Client{Jar: ymm.CookieJar}
}

// func (ymm YoutubeMusicManager) get() error {
