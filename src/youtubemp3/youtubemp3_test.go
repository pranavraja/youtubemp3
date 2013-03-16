package youtubemp3

import (
	"os"
	"strings"
	"testing"
)

func BenchmarkDownloader(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := Video{DownloadUrl: "http://google.com"}
		f, err := os.Create("_tmp.out")
		if err != nil {
			panic(err)
		}
		err = v.Download(f)
		if err != nil {
			panic(err)
		}
	}
	os.Remove("_tmp.out")
}

func BenchmarkGetVideo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v, err := GetVideo("http://www.youtube.com/watch?v=-98o_QOrJho")
		if err != nil {
			panic(err)
		}
		if v.Filename != "The Kooks - Naive.mp3" {
			b.Errorf("Wrong title: %s", v.Filename)
		}
		if !strings.Contains(v.DownloadUrl, "vidtomp3.com/download") {
			b.Errorf("Wrong download url: %s", v.DownloadUrl)
		}
	}
}
