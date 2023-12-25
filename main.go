package main

import (
	"bufio"
	"log"
	"net/http"
	"strings"
)

// "github.com/nk521/go-starry/youtube_music"

func main() {
	headers := `POST /youtubei/v1/browse?key=AIzaSyC9XL3ZjWddXya6X74dJoCTL-WEYFDNX30&prettyPrint=false HTTP/2.0
Host: music.youtube.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:121.0) Gecko/20100101 Firefox/121.0
Accept: */*
Accept-Language: en-US,en;q=0.5
Accept-Encoding: gzip, deflate, br
Referer: https://music.youtube.com/
Content-Type: application/json
X-Goog-Visitor-Id: CgtjM1EzVFBkVFozTSjCuqesBjIKCgJJThIEGgAgXQ%3D%3D
X-Youtube-Bootstrap-Logged-In: true
X-Youtube-Client-Name: 67
X-Youtube-Client-Version: 1.20231214.00.00
Authorization: SAPISIDHASH 1703533915_f477f4d8da139e5b382bcd4a6a974793c549b0cf
X-Goog-AuthUser: 0
X-Origin: https://music.youtube.com
Content-Length: 3593
Origin: https://music.youtube.com
DNT: 1
Connection: keep-alive
Cookie: VISITOR_INFO1_LIVE=c3Q3TPdTZ3M; PREF=f6=40000081&volume=6&f7=140&tz=Asia.Kolkata&repeat=NONE&library_tab_browse_id=FEmusic_liked_playlists&autoplay=true&f5=30000; SID=eQiraGWciVZWj65Mnh1-_Lge6ePLSOENf4lmMhGMFq8rJyuoEj_Fg7XJhFiOOUjbwNucvA.; __Secure-1PSID=eQiraGWciVZWj65Mnh1-_Lge6ePLSOENf4lmMhGMFq8rJyuoSd_6xIBqZ0MfDN7dCfmC-Q.; __Secure-3PSID=eQiraGWciVZWj65Mnh1-_Lge6ePLSOENf4lmMhGMFq8rJyuolTBaAIhtFndbabZo57PonA.; HSID=AvLQ3kMqxi5dpn_Ec; SSID=AQGoc7rCJ1I0EUemz; APISID=LUbMM9KQLId2tz8K/A8kkkUs62_83uGWWx; SAPISID=LIjhl-OsohoxYhVn/AROIR9onzfJC8n5ZH; __Secure-1PAPISID=LIjhl-OsohoxYhVn/AROIR9onzfJC8n5ZH; __Secure-3PAPISID=LIjhl-OsohoxYhVn/AROIR9onzfJC8n5ZH; LOGIN_INFO=AFmmF2swRQIga6GTXEr6Cm-DZ-2H-syTNCrFUdxz2B8838TgPzfQORQCIQCafo0yasl_ZOA1kXA33x5dmOV5uyw3qHMnmVv88jc6kA:QUQ3MjNmejlaOFdMeS0xbHhMdmhvb3ZTN0tsOXlTREdnbUFhMnhxdm5tZzhxUGpnZnlGWHpTSzNPU3g4a2VJOEMyTGVRaHRlX3gxVHEwNUl1c2JwMGRfMlgtMWRUYXY4MHA1ZzBNNGU3RW1YTk1ybGZyd3A4SGl4amo0OVJaY2VfNXdMOW5SWWw4eDhxVzc2U05JWTZxa2gyWlYyMFJOc0V3; SIDCC=ABTWhQHYCDPtMcjYSUZIqLnJAjBKOQOzKg0lxmyhcHxGl_xBt0F4wpd-eIG1PM8rBBQWdcW2bjSA; __Secure-1PSIDCC=ABTWhQF64EqnBk9iupSkUiTIVb22Rx-odEpnfHi39-46_x5R0SBCuf0cavtbAGTacTrRK4NBWDXs; __Secure-3PSIDCC=ABTWhQG04cj_yHTgH7_EkeKsg0IcUCiJm-NJw1cbP34RigTmKeID3Uv2VhegoZhDFLoDaMbNISg; _ga_VCGEPY40VB=GS1.1.1661964189.1.0.1661964191.0.0.0; _ga=GA1.1.1870640887.1661964190; __Secure-1PSIDTS=sidts-CjIBPVxjSi4r232uBeGQL9TFOSfLGHo7lw9pjFjqgShz9QwvK3B8r2qXcUFDpJnTLmeEnhAA; __Secure-3PSIDTS=sidts-CjIBPVxjSi4r232uBeGQL9TFOSfLGHo7lw9pjFjqgShz9QwvK3B8r2qXcUFDpJnTLmeEnhAA; VISITOR_PRIVACY_METADATA=CgJJThIEGgAgXQ%3D%3D; YSC=Vxt28TjLg44
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: same-origin
Sec-Fetch-Site: same-origin
TE: trailers`

	// headers = strings.ReplaceAll(headers, "\n", "\r\n")
	// log.Println(headers)

	// don't forget to make certain the headers end with a second "\r\n"
	reader := bufio.NewReader(strings.NewReader(headers + "\n\n"))
	logReq, err := http.ReadRequest(reader)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(logReq.Header)
	// cmd.Execute()
	// youtube_music.GetCookie()
}
