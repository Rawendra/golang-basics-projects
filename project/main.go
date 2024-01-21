package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("this is init")
	url := "http://google.com"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	bs := make([]byte, 99999)
	response.Body.Read(bs) //reads the response bytes and pushes them in bs
	fmt.Println(string(bs))
}
