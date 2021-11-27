package algorithm

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type point struct {
	x, y float64
}

func (p *point) set_p(rng *rand.Rand) {
	p.x, p.y = rng.Float64(), rng.Float64()
}

func (p *point) in_circle() bool {
	return p.x*p.x+p.y*p.y <= 1
}

func generate(ch chan uint64, n uint64) {
	var buf uint64
	rng := rand.New(rand.NewSource(time.Now().Unix()))
	log.Debug("Started generator")
	p := point{0, 0}
	var i uint64
	for i < n {
		p.set_p(rng)
		if p.in_circle() {
			buf++
		}
		i++
	}
	ch <- buf
}

func MonteCarlo() {
	log.Info("Monte-Carlo-Simulation started")

	// params
	var n_generators, N, inside uint64 = 8, 1e8, 0
	var wg sync.WaitGroup
	ch := make(chan uint64, n_generators)
	for i := 0; i < int(n_generators); i++ {
		wg.Add(1)
		go func() {
			generate(ch, N/n_generators)
			wg.Done()
		}()
	}

	log.Info("All generators started. Waiting for them to return")

	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		inside += i
	}

	log.Info("Simulation done. Showing results")
	fmt.Println("PI: ", 4*inside)
}
