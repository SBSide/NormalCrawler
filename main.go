package main

import (
	"fmt"
	"main/core"
)

func main() {
	/*var results = make(map[string]string)
	urls := []string{
		"https://www.google.com/",
		"https://github.com/",
		"https://discordapp.com/",
		"https://www.amazon.com/",
		"https://www.youtube.com/",
		"https://www.spotify.com/",
		"https://www.instagram.com/",
	}
	for _, url := range urls {
		result := "OK"
		err := core.HitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}*/
	c := make(chan bool)
	people := [2]string{"nyan", "Meow"}
	for _, person := range people {
		go core.IsSexy(person, c)
	}
	fmt.Println(<-c)
	fmt.Println(<-c)
}
