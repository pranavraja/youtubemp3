Download mp3s from youtube URLs, on the command line. Uses an (undocumented) API at vidtomp3.com to do the conversion

`go get github.com/pranavraja/youtubemp3`

## Usage

Assuming `$GOPATH/bin` is in your `$PATH`:

    go build
    youtubemp3

Wait for URLs from `stdin` and downloads them to local mp3 files as they arrive, with format `VideoTitle.mp3`. Print each filename to `stdout` when complete. Print errors to `stderr`.

## Other examples

    youtubemp3 -i youtube-urls.txt

Download tracks from URLs defined in `youtube-urls.txt`, to local files with format `VideoTitle.mp3`.

    youtubemp3 -p http://www.youtube.com/playlist?list=PL307A9B1AC2C63F98

Download all tracks from a given playlist, to local files with format `VideoTitle.mp3`
