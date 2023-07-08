// example code for rate limiting

package main

import (
	"fmt"
	"time"
)

func main() {
	// limit to 5 per second
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// limiter channel will receive a value every 200ms.
	limiter := time.Tick(200 * time.Millisecond)

	// block on a receive from the limiter before serving each request.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// allow short bursts of requests in our rate limiting scheme while
	// preserving the overall rate limit.
	// this is done by buffering our limiter channel.
	// the burstyLimiter channel will allow bursts of up to 3 events.
	burstyLimiter := make(chan time.Time, 3)

	// fill up the channel to represent allowed bursting.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// every 200ms we'll try to add a new value to burstyLimiter, up to its limit of 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// simulate 5 more incoming requests.
	// the first 3 of these will benefit from the burst capability of burstyLimiter.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	// block on a receive from burstyLimiter before serving each request.
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
