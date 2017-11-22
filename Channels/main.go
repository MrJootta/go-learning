package main

import (
	"net/http"
	"fmt"
	"time"
)

func main() {
	links := []string{
		"http://id.atlastock.com",
		"http://pt.atlastock.com",
		"http://ro.atlastock.com",
		"http://ua.atlastock.com",
		"http://pl.atlastock.com",
	}

	// Create channel
	c := make(chan string)


	for _, link := range links {
		go executeRequest(link, c)
	}

	// Need same number of print as links / channels are created
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)

	//Smart way to do it
	//for i := 0; i < len(links); i++ {
		//fmt.Println(<-c)
	//}

	// Infinite loop
	//for {
	//		go executeRequest(<-c, c)
	//}

	//Alternative to infinite loop
	for l := range c {
		//Literal function to add sleep we pass the right link as a param
		go func(link string) {
			time.Sleep(time.Second)
			executeRequest(link, c)
		}(l)
	}
}

func executeRequest(link string, c chan string){
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " Is Down!!!")
		c <- link
		return
	}

	fmt.Println(link, " Is UP :)")
	c <- link
}
