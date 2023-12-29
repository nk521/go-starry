package youtube_music

import (
	"crypto/sha1"
	"encoding/hex"
	"net/http"
	"strconv"
	"time"
)

func sapisidFromCookie(cookie string) string {
	header := http.Header{}
	header.Add("Cookie", cookie)
	request := http.Request{Header: header}

	cookieval, _ := request.Cookie("__Secure-3PAPISID")
	return cookieval.Value
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

func getAuthorization(auth string) string {
	h := sha1.New()

	unix_timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	final := unix_timestamp + " " + auth
	h.Write([]byte(final))

	return "SAPISIDHASH " + unix_timestamp + "_" + hex.EncodeToString(h.Sum(nil))
}
