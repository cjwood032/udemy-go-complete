package main

import (
	"fmt"
	"net/http"
)

func getStatus(c chan string, link string) {
	_, err := http.Get(link)
	if err != nil {
		//fmt.Println(link, " may be down")
		c <- link + " might be down"
	} else {
		//fmt.Println(link, "  is up", link)
		c <- link + " looks like it's up"
	}
}

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go getStatus(c, link)

	}
	for i := 0; i < len(links)-1; i++ {
		fmt.Println(<-c)
	}
	fmt.Println(<-c)
}
