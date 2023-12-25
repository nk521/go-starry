package youtube_music

import (
	"errors"
	"net/http"
	"net/url"
)

func sapisid_from_cookie(client *http.Client) (string, error) {
	_url, _ := url.Parse("https://music.youtube.com")
	for _, cookie := range client.Jar.Cookies(_url) {
		if cookie.Name == "__Secure-3PAPISID" {
			return cookie.Value, nil
		}
	}

	return "", errors.New("couldn't find sapisid cookie in browser cookies! please run `starry cookie` again")

}
