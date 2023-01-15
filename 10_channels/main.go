package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://amazon.com",
		"http://stackoverflow.com",
		"http://google.com",
		"http://facebook.com",
	}

	// channel declaration
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c) // declare child go routine for each function
	}

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// infinite loop
	// for {
	// 	go checkLink(<-c, c)
	// }
	// alt syntax for more readibility
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(s string, c chan string) {
	_, err := http.Get(s)
	if err != nil {
		fmt.Println(s, "might be down!")
		// c <- "Might be down I think"
		c <- s
		return
	}

	fmt.Println(s, "is up!")
	// c <- "Yep its up"
	c <- s
}
