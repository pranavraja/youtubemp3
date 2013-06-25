package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/pranavraja/youtubemp3/video"
	"io"
	"os"
	"sync"
)

var inputFile string
var playlist string

func init() {
	flag.StringVar(&inputFile, "i", "-", "Input file (default stdin)")
	flag.StringVar(&playlist, "p", "", "Playlist to download videos from")
}

func download(vid *video.Video) {
	file, err := os.Create(vid.Filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	defer file.Close()
	vid.Download(file)
	println(vid.Filename)
}

func fetchAllFromPlaylist(playlistUrl string) {
	if playlistUrl == "" {
		return
	}
	videos, err := video.GetPlaylist(playlistUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	var wg sync.WaitGroup
	for _, vid := range videos {
		wg.Add(1)
		go func(v *video.Video) {
			download(v)
			wg.Done()
		}(vid)
	}
	wg.Wait()
}

func fetchAll(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	var wg sync.WaitGroup
	for scanner.Scan() {
		url := scanner.Text()
		if url == "" {
			return
		}
		vid, err := video.GetVideo(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return
		}
		wg.Add(1)
		go func(v *video.Video) {
			download(v)
			wg.Done()
		}(vid)
	}
	wg.Wait()
}

func main() {
	flag.Parse()
	fetchAllFromPlaylist(playlist)
	fileReader := os.Stdin
	if inputFile != "-" {
		var err error
		fileReader, err = os.Open(inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return
		}
	}
	fetchAll(fileReader)
}
