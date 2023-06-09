package main

import (
	"fmt"
)

var messages = []string{
	"The world itself's",
	"just one big hoax.",
	"Spamming each other with our",
	"running commentary of bullshit,",
	"masquerading as insight, our social media",
	"faking as intimacy.",
	"Or is it that we voted for this?",
	"Not with our rigged elections,",
	"but with our things, our property, our money.",
	"I'm not saying anything new.",
	"We all know why we do this,",
	"not because Hunger Games",
	"books make us happy,",
	"but because we wanna be sedated.",
	"Because it's painful not to pretend,",
	"because we're cowards.",
	"- Elliot Alderson",
	"Mr. Robot",
}

func producer(dlink chan<- string) {
	for _, m := range messages {
		dlink <- m
	}
	close(dlink)
}

func consumer(dlink <-chan string, ctldone chan<- bool) {
	for b := range dlink {
		fmt.Println(b)
	}
	ctldone <- true
}

func main() {
	link := make(chan string)
	done := make(chan bool)
	go producer(link)
	go consumer(link, done)
	<-done
}
