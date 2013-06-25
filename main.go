package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/pranavraja/youtubemp3/video"
	"os"
	"strings"
	"sync"
)

var inputFile string
var playlist string

func init() {
	flag.StringVar(&inputFile, "i", "-", "Input file (default stdin)")
	flag.StringVar(&playlist, "p", "", "Playlist to download videos from")
}

func doForEachLineInFile(fileName string, handler func(string)) (err error) {
	fileReader := os.Stdin
	if fileName != "-" {
		fileReader, err = os.Open(fileName)
		if err != nil {
			return err
		}
	}
	buffered := bufio.NewReader(fileReader)
	var wg sync.WaitGroup
	for {
		line, err := buffered.ReadString('\n')
		if err != nil {
			break
		}
		wg.Add(1)
		go func(line string) {
			handler(line)
			wg.Done()
		}(line)
	}
	wg.Wait()
	return err
}

func download(video *video.Video) {
	file, err := os.Create(video.Filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	defer file.Close()
	video.Download(file)
	println(video.Filename)
}

func main() {
	flag.Parse()
	if playlist != "" {
		videos, err := video.GetPlaylist(playlist)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return
		}
		for _, video := range videos {
			download(video)
		}
		return
	}
	doForEachLineInFile(inputFile, func(videoUrl string) {
		url := strings.TrimRight(videoUrl, "\n")
		if url == "" {
			return
		}
		video, err := video.GetVideo(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return
		}
		download(video)
	})
}
