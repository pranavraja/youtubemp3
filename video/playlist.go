package video

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
)

type playlistFeed struct {
	Feed struct {
		Entry []struct {
			Link []struct {
				Href string
			}
		}
	}
}

func GetPlaylist(url string) ([]*Video, error) {
	playlistIdFromUrl := regexp.MustCompile(`list=(\w+)`)
	playlistId := playlistIdFromUrl.FindStringSubmatch(url)
	if playlistId == nil {
		return nil, errors.New("url could not parsed")
	}
	resp, err := http.Get("https://gdata.youtube.com/feeds/api/playlists/" + playlistId[1] + "?v=2&alt=json")
	if err != nil {
		return nil, err
	}
	feed := new(playlistFeed)
	err = json.NewDecoder(resp.Body).Decode(feed)
	if err != nil {
		return nil, err
	}
	videos := make([]*Video, 0, len(feed.Feed.Entry))
	for _, entry := range feed.Feed.Entry {
		url := entry.Link[0].Href
		video, err := GetVideo(url)
		if err != nil {
			continue
		}
		videos = append(videos, video)
	}
	return videos, nil
}
