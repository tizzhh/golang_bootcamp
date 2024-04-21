package main

import (
	"fmt"
	"presentHeap/present_heap"
)

func main() {
	prArr := []presentHeap.Present{{Value: 5, Size: 1}, {Value: 4, Size: 5}, {Value: 3, Size: 1}, {Value: 5, Size: 2}}
	fmt.Println(presentHeap.GetNCoolestPresents(prArr, 5))
}
