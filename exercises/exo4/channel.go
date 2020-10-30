package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Transform this code, so that we only wait for the first request to finish

func main() {
	rand.Seed(time.Now().UnixNano())
	myChan := make(chan string)
	res1 := fakeHttpRequest("google.com")
	go func() {
		defer close(myChan)
		myChan <- fakeHttpRequest("bing.com")
	}()
	res2 := <-myChan
	fmt.Printf(res1)
	fmt.Printf(res2)
}

func fakeHttpRequest(url string) string {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) // simulate network latency
	return fmt.Sprintf("got it from %s", url)
}
