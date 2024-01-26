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

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for index := 0; index < len(links); index++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "the link is not working it seems")
		c <- "the link is not working it seems"
		return
	}
	fmt.Println(link, "link is working fine")
	c <- "link is working fine"

}
