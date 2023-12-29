package youtube_music

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"

	"strings"

	"github.com/nk521/go-starry/config"
	log "github.com/nk521/go-starry/log"
	"github.com/nk521/go-starry/util"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
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
	Context       map[string]interface{}
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

	req.Header = ymm.Headers.Clone()

	resp, err := ymm.Client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	return resp
}

func (ymm YoutubeMusicManager) sendPOSTRequest(endpoint string, body map[string]interface{}, additional_params string) []byte {
	params := YTM_PARAMS + YTM_PARAMS_KEY
	_ = params

	maps.Copy(body, ymm.Context)

	jsonStr, err := json.Marshal(body)
	if err != nil {
		log.Panicln(err)
	}
	req, err := http.NewRequest(http.MethodPost, YTM_BASE_API+endpoint+params+additional_params, strings.NewReader(string(jsonStr)))
	if err != nil {
		log.Panicln(err)
	}

	req.Header = ymm.Headers.Clone()
	origin := req.Header.Get("origin")
	if origin == "" {
		origin = req.Header.Get("x-origin")
	}
	req.Header.Set("authorization", getAuthorization(ymm.sapisid+" "+origin))

	req.AddCookie(&http.Cookie{Name: "SOCS", Value: "CAI"})

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

func NewYoutubeMusicManager() *YoutubeMusicManager {
	ymm := YoutubeMusicManager{}
	ymm.Headers = &http.Header{}
	// set headers
	to_delete := []string{"host", "content-length", "accept-encoding"}
	_headers := util.ParseHeadersFromString(config.GetConfig().Login.Headers)

	for k, v := range *_headers {
		_k := strings.ToLower(k)
		if !(strings.HasPrefix(_k, "sec") && slices.Contains(to_delete, _k)) {
			ymm.Headers.Set(k, v[0])
		}

	}

	ymm.Headers.Set("user-agent", USER_AGENT)
	ymm.Headers.Set("accept", "*/*")
	ymm.Headers.Set("accept-encoding", "gzip, deflate")
	ymm.Headers.Set("content-type", "application/json")
	ymm.Headers.Set("content-encoding", "gzip")
	ymm.Headers.Set("origin", YTM_DOMAIN)

	cookie := ymm.Headers.Get("Cookie")
	ymm.sapisid = sapisidFromCookie(cookie)

	// init context
	ymm.Context = initContext()

	// and finally, the client
	ymm.Client = &http.Client{}

	return &ymm
}

func (ymm YoutubeMusicManager) GetHomePage(limit int) {
	endpoint := "browse"
	body := map[string]interface{}{"browseId": "FEmusic_home"}
	response := ymm.sendPOSTRequest(endpoint, body, "")

	_ = response
}
