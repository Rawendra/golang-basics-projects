package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.facebook.com/",
		"https://www.primevideo.com",
		"https://www.youtube.com"}

	for _, link := range links {
		checkLink(link)
	}

}

func checkLink(link string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "the link is not working it seems")
		return
	}
	fmt.Println(link, "link is working fine")

}
