package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func multiplex(chans ...chan interface{}) chan interface{} {
	chanRes := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(len(chans))
	for _, ch := range chans {
		go func(ch chan interface{}) {
			defer wg.Done()
			for val := range ch {
				chanRes <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(chanRes)
	}()

	return chanRes
}

func main() {
	randChans := make([]chan interface{}, 10)
	rand.Shuffle(len(randChans), func(i, j int) {
		randChans[i], randChans[j] = randChans[j], randChans[i]
	})
	for i := 0; i < len(randChans); i++ {
		randChans[i] = make(chan interface{})
		go func(ch chan interface{}, i int) {
			ch <- i
			close(ch)
		}(randChans[i], i)
	}
	res := multiplex(randChans...)
	for val := range res {
		fmt.Println(val)
	}
}
