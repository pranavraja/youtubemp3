package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"youtubemp3"
)

var inputFile string

func init() {
	flag.StringVar(&inputFile, "i", "-", "Input file (default stdin)")
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
	for {
		line, err := buffered.ReadString('\n')
		if err != nil {
			break
		}
		handler(line)
	}
	return err
}

func main() {
	flag.Parse()
	doForEachLineInFile(inputFile, func(youtubeUrl string) {
		video, err := youtubemp3.GetVideo(strings.TrimRight(youtubeUrl, "\n"))
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		file, err := os.Create(video.Filename)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		defer file.Close()
		video.Download(file)
		println(video.Filename)
	})
}