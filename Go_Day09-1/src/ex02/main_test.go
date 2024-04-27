package main

import (
	"math/rand"
	"testing"
)

const (
	NUM_OF_CHANS int = 10
)

func TestMain(t *testing.T) {
	randChans := make([]chan interface{}, NUM_OF_CHANS)
	testIndexes := make([]interface{}, 0, len(randChans))

	for i := 0; i < len(randChans); i++ {
		testIndexes = append(testIndexes, rand.Intn(69))
	}

	for i, val := range testIndexes {
		randChans[i] = make(chan interface{})
		elem := val
		go func(ch chan interface{}, i interface{}) {
			ch <- elem
			close(ch)
		}(randChans[i], elem)
	}
	res := multiplex(randChans...)

	results := make(map[interface{}]bool)

	for val := range res {
		results[val] = true
	}

	for _, val := range testIndexes {
		if !results[val] {
			t.Fatalf("Test data %v not recieved\n", val)
		}
	}
}
