Download mp3s from youtube URLs, on the command line. Uses an (undocumented) API at vidtomp3.com to do the conversion

## Usage

    go build
    ./youtubemp3 -h

Waits for URLs from `stdin` and downloads them to local mp3 files as they arrive, with format `VideoTitle.mp3`. Prints each filename after it is downloaded.

## Other examples

    ./youtubemp3 -i youtube-urls.txt

Downloads youtube URLs, one per line in `youtube-urls.txt`, to local files with format `VideoTitle.mp3`.
