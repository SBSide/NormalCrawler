package main

import (
	"fmt"
	"main/core"
)

func main() {
	results := make(map[string]string)
	c := make(chan core.RequestResult)
	urls := []string{
		"https://www.google.com/",
		"https://github.com/",
		"https://discordapp.com/",
		"https://www.amazon.com/",
		"https://www.youtube.com/",
		"https://www.spotify.com/",
		"https://www.instagram.com/",
		"https://www.youtube.com/",
		"https://last.fm/",
		"https://www.reddit.com/",
	}
	for _, url := range urls {
		go core.HitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.Url] = result.Status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}
}
