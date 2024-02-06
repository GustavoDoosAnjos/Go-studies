package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	sites := []string{
		"http://www.google.com",
		"http://www.yahoo.com",
		"http://www.youtube.com",
		"http://www.udemy.com",
	}

	c := make(chan string)

	for _, site := range sites {
		go makeRequest(site, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(2 * time.Second)
			makeRequest(link, c)
		}(l)
	}
}

func makeRequest(site string, c chan string) {
	_, err := http.Get(site)
	if err != nil {
		fmt.Printf("An error ocurred on %v \n", site)
		c <- site
	}

	fmt.Printf("The site %v is OK \n", site)
	c <- site
}
