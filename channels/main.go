package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	URLS := []string{
		"google.com",
		"golang.org",
		"stackoverflow.com",
		"facebook.com",
		"amazon.com",
	}

	//Channels are ways go-routines communicate/send messages with each other.
	//A channel has type chan and then a type of what kind of message is passed,
	//in this case, a string
	c := make(chan string)

	//By default, Go only uses 1 core. So when you have multiple
	//go-routines, only 1 is ran to completion or until a blocking
	//call is made. Can change it to use multiple cores and run
	//multiple routines at once.
	for _, link := range URLS {
		go checkLink(link, c)
	}

	//Continuously check the links using an endless loop by waiting
	//on the channel to return a value, then going into the for loop.
	//Pass in the link we put in the channel back into the checkLink
	//function along with the channel itself
	for l := range c {
		//Uses a lambda (known as a function literal) in the go-routine so
		//we can have a wait for every request after the initial one.
		go func(link string) {
			time.Sleep(time.Second * 5)
			checkLink(link, c)
		}(l) //Passing l from loop into the function so it is copied
		//(since its by value) and thus does not reference the same
		//memory address as others and becomes thread-safe. Don't want
		//to access the same value by different go-routines.
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get("http://" + link)
	if err != nil {
		fmt.Println(link, "is down")
		c <- link
	} else {
		fmt.Println(link, "is up")
		c <- link
	}
}
