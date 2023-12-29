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
	"github.com/nk521/go-starry/util"
)

var (
	YTM_DOMAIN     = "https://music.youtube.com"
	YTM_BASE_API   = YTM_DOMAIN + "/youtubei/v1/"
	YTM_PARAMS     = "?alt=json"
	YTM_PARAMS_KEY = "&key=AIzaSyC9XL3ZjWddXya6X74dJoCTL-WEYFDNX30"
	USER_AGENT     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:120.0) Gecko/20100101 Firefox/120.0"
)

type YoutubeMusicManager struct {
	Client        *http.Client
	CookieJar     *cookiejar.Jar
	HeaderCookies string
	Headers       *http.Header
	Context       map[string]map[string]map[string]string
	sapisid       string
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

func (ymm YoutubeMusicManager) Init() {
	// set headers
	ymm.Headers = util.ParseHeadersFromString(config.GetConfig().Login.Headers)

	cookie := ymm.Headers.Get("Cookie")
	ymm.sapisid = sapisidFromCookie(cookie)

	// make cookie jar
	{
		_url, _ := url.Parse(YTM_DOMAIN)
		_cookie_jar, err := cookiejar.New(nil)
		if err != nil {
			log.Panicln(err)
		}
		ymm.CookieJar = _cookie_jar
		ymm.CookieJar.SetCookies(_url, []*http.Cookie{{Name: "SOCS", Value: "CAI"}})
	}

	// init context
	ymm.Context = initContext()
	ymm.Context["context"]["client"]["hl"] = "en"

	// and finally, the client
	ymm.Client = &http.Client{Jar: ymm.CookieJar}
}

// func (ymm YoutubeMusicManager) getVisitorId() []byte {
// 	resp := ymm.sendGETRequest(YTM_DOMAIN, nil)
// 	defer resp.Body.Close()
// 	response_body, _ := io.ReadAll(resp.Body)

// 	return response_body
// }
