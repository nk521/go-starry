package youtube_music

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func sapisidFromCookie(client *http.Client) (string, error) {
	_url, _ := url.Parse("https://music.youtube.com")
	for _, cookie := range client.Jar.Cookies(_url) {
		if cookie.Name == "__Secure-3PAPISID" {
			return cookie.Value, nil
		}
	}

	return "", errors.New("couldn't find sapisid cookie in browser cookies! please run `starry cookie` again")

}

func initContext() map[string]map[string]map[string]string {
	return map[string]map[string]map[string]string{
		"context": {
			"client": {
				"clientName":    "WEB_REMIX",
				"clientVersion": "1." + ".01.00",
			},
			"user": {},
		},
	}
}

// func getVisitorId

func getAuthorization(auth string) string {
	h := sha1.New()

	unix_timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	final := unix_timestamp + " " + auth
	h.Write([]byte(final))

	return "SAPISIDHASH " + unix_timestamp + "_" + hex.EncodeToString(h.Sum(nil))
}
