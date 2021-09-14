package main

import (
	"log"
	"time"

	"github.com/Frosin/limiter/limiter"
)

func main() {
	testLimiterOpts := limiter.LimiterParams{
		TimeInterval: time.Millisecond * 50,
		CountLimit:   2,
	}
	mapLimiter := limiter.NewMapLimiter()

	log.Printf("limiter options: timeInterval: %v, count limit: %d\n",
		testLimiterOpts.TimeInterval,
		testLimiterOpts.CountLimit)
	for i := 1; i < 10; i++ {
		time.Sleep(time.Millisecond * 10)
		ok, err := mapLimiter.Check("test", time.Now(), testLimiterOpts)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("request %d, result=%v\n", i, ok)
	}
}
