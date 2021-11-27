package main

import (
	"fmt"
	"math/rand"
	"montecarlo/algorithm"
	"time"

	log "github.com/sirupsen/logrus"
)

func benchmark(fn func(), n int) time.Duration {
	start := time.Now()
	for i := 0; i < n; i++ {
		fn()
	}
	return time.Since(start)
}

func main() {
	rand.Seed(time.Now().Unix())
	log.SetLevel(log.InfoLevel)
	duration := benchmark(algorithm.MonteCarlo, 1)
	fmt.Printf("Elapsed Time: %s\n", duration)
}
