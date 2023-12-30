package youtube_music

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"regexp"
	"strconv"

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

// func (ymm YoutubeMusicManager) sendGETRequest(url string, params map[string]string) *http.Response {
// 	req, err := http.NewRequest(http.MethodGet, url, nil)
// 	if err != nil {
// 		log.Panicln(err)
// 	}

// 	q := req.URL.Query()
// 	for k, v := range params {
// 		q.Add(k, v)
// 	}
// 	req.URL.RawQuery = q.Encode()

// 	req.Header = ymm.Headers.Clone()

// 	resp, err := ymm.Client.Do(req)
// 	if err != nil {
// 		log.Panicln(err)
// 	}
// 	return resp
// }

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
	sapi := getAuthorization(sapisidFromCookie(ymm.Headers.Get("Cookie")) + " " + origin)
	req.Header.Set("authorization", sapi)
	log.Println(sapi)

	req.AddCookie(&http.Cookie{Name: "SOCS", Value: "CAI"})

	resp, err := ymm.Client.Do(req)
	if err != nil {
		log.Panicln(err)
	}

	if resp.StatusCode >= 400 {
		log.Panicln(fmt.Errorf("got %s @ %s", resp.Status, req.URL.String()))
	}

	defer resp.Body.Close()
	// response_body, _ := io.ReadAll(resp.Body)
	// s, _ := json.Marshal(response_body)

	// var dst []byte
	// base64.StdEncoding.Decode(dst, response_body)
	gr, err := gzip.NewReader(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	defer gr.Close()

	result, _ := io.ReadAll(gr)

	return result
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

func (ymm YoutubeMusicManager) parseWatchPlaylist(data *MusicTwoRowItemRenderer) map[string]interface{} {
	return map[string]interface{}{
		"title":      data.Title.Runs[0].Text,
		"playlistId": data.NavigationEndpoint.WatchPlaylistEndpoint.PlaylistID,
	}
}

func (ymm YoutubeMusicManager) parseSongRuns(runs []SubtitleRun) map[string]interface{} {
	parsed := map[string][]map[string]string{
		"artists": {},
	}
	_parsed := map[string]interface{}{}

	for i, run := range runs {
		if i%2 != 0 {
			continue
		}
		text := run.Text
		if run.NavigationEndpoint != nil {
			item := map[string]string{
				"name": text,
				"id":   run.NavigationEndpoint.BrowseEndpoint.BrowseID,
			}

			if len(item["id"]) > 0 && (strings.HasPrefix(item["id"], "MPRE") || strings.Contains(item["id"], "release_detail")) {
				_parsed["album"] = item
			} else {
				parsed["artists"] = append(parsed["artists"], item)
			}
		} else {
			if matched, _ := regexp.MatchString(`^\d([^ ])* [^ ]*$`, text); matched && i > 0 {
				_parsed["views"] = strings.Split(text, " ")[0]
			} else if matched, _ := regexp.MatchString(`^(\d+:)*\d+:\d+$`, text); matched {
				_parsed["duration"] = text

				duration_split := []int{}
				for _, x := range strings.Split(text, ":") {
					t, _ := strconv.Atoi(x)
					duration_split = append(duration_split, t)
				}
				duration_split[0] = duration_split[0] * 3600
				duration_split[1] = duration_split[1] * 60
				duration_split[2] = duration_split[2] * 1

				sum := 0
				for i := range duration_split {
					sum += duration_split[i]
				}

				_parsed["duration_seconds"] = sum
			} else if matched, _ := regexp.MatchString(`^\d{4}$`, text); matched {
				_parsed["year"] = text
			} else {
				parsed["artists"] = append(parsed["artists"], map[string]string{
					"name": text,
					"id":   "",
				})
			}
		}
	}

	for k, v := range parsed {
		_parsed[k] = v
	}

	return _parsed

}

func (ymm YoutubeMusicManager) parseSong(data *MusicTwoRowItemRenderer) map[string]interface{} {
	song := map[string]interface{}{
		"title":      data.Title.Runs[0].Text,
		"videoId":    data.NavigationEndpoint.WatchEndpoint.VideoID,
		"playlistId": data.NavigationEndpoint.WatchEndpoint.PlaylistID,
	}

	maps.Copy(song, ymm.parseSongRuns(data.Subtitle.Runs))

	return song
}

func (ymm YoutubeMusicManager) parseAlbum(data *MusicTwoRowItemRenderer) map[string]interface{} {
	isExplicit := false
	if len(data.SubtitleBadges) > 0 && data.SubtitleBadges[0].MusicInlineBadgeRenderer.AccessibilityData.AccessibilityData.Label != "" {
		isExplicit = true
	}
	return map[string]interface{}{
		"title":      data.Title.Runs[0].Text,
		"year":       data.Subtitle.Runs[2].Text,
		"browseId":   data.Title.Runs[0].NavigationEndpoint.BrowseEndpoint.BrowseID,
		"isExplicit": isExplicit,
	}

}

func (ymm YoutubeMusicManager) parseRelatedArtist(data *MusicTwoRowItemRenderer) map[string]interface{} {
	subs := data.Subtitle.Runs[0].Text

	if subs != "" {
		subs = strings.Split(subs, " ")[0]
	}

	return map[string]interface{}{
		"title":       data.Title.Runs[0].Text,
		"browseId":    data.Title.Runs[0].NavigationEndpoint.BrowseEndpoint.BrowseID,
		"subscribers": subs,
	}
}

func (ymm YoutubeMusicManager) parsePlaylist(data *MusicTwoRowItemRenderer) map[string]interface{} {
	playlist := map[string]interface{}{
		"title":      data.Title.Runs[0].Text,
		"playlistId": data.Title.Runs[0].NavigationEndpoint.BrowseEndpoint.BrowseID[2:],
	}

	subtitle := data.Subtitle

	if len(subtitle.Runs) > 0 {
		_sub := []string{}
		for _, run := range subtitle.Runs {
			_sub = append(_sub, run.Text)
		}
		playlist["description"] = strings.Join(_sub, "")

		if len(subtitle.Runs) == 3 {
			if matched, _ := regexp.MatchString(`\d+ `, data.Subtitle.Runs[2].Text); matched {
				playlist["count"] = strings.Split(data.Subtitle.Runs[2].Text, " ")[0]
				artists := []map[string]string{}

				runs := subtitle.Runs[:1]
				for j := 0; j < int(len(runs)/2)+1; j++ {
					artists = append(artists, map[string]string{
						"name": runs[j*2].Text,
						"id":   runs[j*2].NavigationEndpoint.BrowseEndpoint.BrowseID,
					})
				}

				playlist["author"] = artists
			}
		}
	}

	return playlist

}

func (ymm YoutubeMusicManager) parseSongArtists(data *MusicResponsiveListItemRenderer, index int) []map[string]string {
	_flex := ymm.getFlexColumnItem(data, index)
	if _flex == nil {
		return nil
	}
	flex := *_flex
	runs := flex.Text.Runs
	artists := []map[string]string{}

	for j := 0; j < int(len(runs)/2)+1; j++ {
		id := ""
		if runs[j*2].NavigationEndpoint != nil {
			id = runs[j*2].NavigationEndpoint.BrowseEndpoint.BrowseID
		}
		artists = append(artists, map[string]string{
			"name": runs[j*2].Text,
			"id":   id,
		})
	}

	return artists
}

func (ymm YoutubeMusicManager) parseSongFlat(data *MusicResponsiveListItemRenderer) map[string]interface{} {
	cols := []*MusicResponsiveListItemFlexColumnRenderer{}
	for col := 0; col < len(data.FlexColumns); col++ {
		temp := ymm.getFlexColumnItem(data, col)
		cols = append(cols, temp)
	}

	isExplicit := false
	if len(data.Badges) > 0 && data.Badges[0].MusicInlineBadgeRenderer.AccessibilityData.AccessibilityData.Label != "" {
		isExplicit = true
	}

	song := map[string]interface{}{
		"title":      cols[0].Text.Runs[0].Text,
		"videoId":    cols[0].Text.Runs[0].NavigationEndpoint.WatchEndpoint.VideoID,
		"artists":    ymm.parseSongArtists(data, 1),
		"isExplicit": isExplicit,
	}

	if len(cols) > 2 && cols[2] != nil && cols[2].Text.Runs[0].NavigationEndpoint != nil {
		song["album"] = map[string]string{
			"name": cols[2].Text.Runs[0].Text,
			"id":   cols[2].Text.Runs[0].NavigationEndpoint.BrowseEndpoint.BrowseID,
		}

	} else {
		runs := cols[1].Text.Runs
		song["views"] = strings.Split(runs[len(runs)-1].Text, " ")[0]
	}

	return song
}

func (ymm YoutubeMusicManager) getFlexColumnItem(data *MusicResponsiveListItemRenderer, index int) *MusicResponsiveListItemFlexColumnRenderer {
	if len(data.FlexColumns) <= index {
		return nil
	}
	temp := data.FlexColumns[index].MusicResponsiveListItemFlexColumnRenderer
	if len(temp.Text.Runs) > 0 {
		return &temp
	}

	return &temp
}

func (ymm YoutubeMusicManager) parseMixedContent(rows []SectionListRendererContent) []map[string]interface{} {
	items := []map[string]interface{}{}
	var title string
	var contents interface{}
	// x := MusicDescriptionShelfRenderer{}

	for _, row := range rows {
		if len(row.MusicDescriptionShelfRenderer.Header.Runs) > 0 {
			results := row.MusicDescriptionShelfRenderer
			title = results.Header.Runs[0].Text
			contents = results.Description.Runs[0].Text
		} else {
			if len(row.MusicCarouselShelfRenderer.Contents) == 0 {
				continue
			}
			title = row.MusicCarouselShelfRenderer.Header.MusicCarouselShelfBasicHeaderRenderer.Title.Runs[0].Text
			log.Println(title)
			_contents := []map[string]interface{}{}
			for _, content := range row.MusicCarouselShelfRenderer.Contents {
				_data := content.MusicTwoRowItemRenderer
				_content := map[string]interface{}{}
				if _data != nil {
					temp := _data.Title.Runs[0].NavigationEndpoint
					var page_type PageType
					if temp == nil {
						page_type = PageType("")
					} else {
						page_type = temp.BrowseEndpoint.BrowseEndpointContextSupportedConfigs.BrowseEndpointContextMusicConfig.PageType
					}

					switch page_type {
					case MusicPageTypeAlbum:
						_content = ymm.parseAlbum(_data)

					case MusicPageTypeArtist:
						_content = ymm.parseRelatedArtist(_data)

					case MusicPageTypePlaylist:
						_content = ymm.parsePlaylist(_data)

					default:
						temp := _data.NavigationEndpoint.WatchPlaylistEndpoint
						if temp != nil && temp.PlaylistID != "" {
							_content = ymm.parseWatchPlaylist(_data)
						} else {
							_content = ymm.parseSong(_data)
						}
					}

				} else {
					_data := content.MusicResponsiveListItemRenderer
					_content = ymm.parseSongFlat(_data)
				}

				_contents = append(_contents, _content)
			}
			contents = _contents
			fmt.Println(title)
		}

		items = append(items, map[string]interface{}{"title": title, "contents": contents})
	}

	return items
}

func (ymm YoutubeMusicManager) GetHomePage(limit int) {
	endpoint := "browse"
	body := map[string]interface{}{"browseId": "FEmusic_home"}
	response := ymm.sendPOSTRequest(endpoint, body, "")

	st := SchemaGetHomePage{}
	// var st map[string]interface{}
	err := json.Unmarshal(response, &st)
	if err != nil {
		log.Panicln(err)
	}

	results := st.Contents.SingleColumnBrowseResultsRenderer.Tabs[0].TabRenderer.Content.SectionListRenderer.Contents
	home := ymm.parseMixedContent(results)
	s, _ := json.MarshalIndent(home, "", ` `)
	fmt.Println(string(s))
	// TODO: change response from objects to strings
	// home := []
	// SchemaGetHomePage
}
