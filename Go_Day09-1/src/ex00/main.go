package main

import (
	"fmt"
	"time"
)

func sleepSort(slc []int) chan int {
	ch := make(chan int)
	for _, val := range slc {
		go func(e int) {
			time.Sleep(time.Duration(e) * time.Second)
			ch <- e
		}(val)
	}
	return ch
}

func main() {
	arr := []int{5, 4, 3, 1, 2}
	ch := sleepSort(arr)
	defer close(ch)
	for i := 0; i < len(arr); i++ {
		fmt.Println(<-ch)
	}
}
