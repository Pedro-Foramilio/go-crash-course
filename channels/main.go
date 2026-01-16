package main

import (
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://golang.org",
		"http://stackoverflow.com",
		"http://google.com",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string, chann chan string) {
			time.Sleep(5 * time.Second)
			checkLink(link, chann)
		}(l, c)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		println(link, "might be down!")
		c <- link
		return
	}
	println(link, "is up!")
	c <- link
}
