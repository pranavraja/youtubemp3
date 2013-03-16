package youtubemp3

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Video struct {
	DownloadUrl string
	Filename    string
}

func (video *Video) Download(output io.Writer) error {
	resp, err := http.Get(video.DownloadUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(output, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func GetVideo(youtubeUrl string) (video *Video, err error) {
	postResponse, err := http.PostForm("http://www.vidtomp3.com/cc/conversioncloud.php", url.Values{"mediaurl": {youtubeUrl}})
	if err != nil {
		return
	}
	defer postResponse.Body.Close()
	body, err := ioutil.ReadAll(postResponse.Body)
	if err != nil {
		return
	}
	var responseJson map[string]interface{}
	// Yeah so vidtomp3 returns json wrapped in parentheses, e.g. ({"response": "success"}), so we have to get rid of that
	err = json.Unmarshal(RemovePrefixAndSuffixParentheses(body), &responseJson)
	if err != nil {
		return
	}
	statusUrl := responseJson["statusurl"].(string) + "&json"
	statusResponse, err := http.Get(statusUrl)
	if err != nil {
		return
	}
	defer statusResponse.Body.Close()
	body, err = ioutil.ReadAll(statusResponse.Body)
	var statusJson map[string]interface{}
	err = json.Unmarshal(RemovePrefixAndSuffixParentheses(body), &statusJson)
	if err != nil {
		return
	}
	video = &Video{DownloadUrl: statusJson["downloadurl"].(string), Filename: statusJson["file"].(string)}
	return
}

func RemovePrefixAndSuffixParentheses(parenthesisedString []byte) []byte {
	return []byte(strings.TrimLeft(strings.TrimRight(string(parenthesisedString), ") "), " ("))
}
