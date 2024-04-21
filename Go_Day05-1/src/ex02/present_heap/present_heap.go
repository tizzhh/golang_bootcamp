package presentHeap

import (
	"container/heap"
	"fmt"
)

type Present struct {
	Value, Size int
}

type PresentHeap []Present

func (ph PresentHeap) Len() int {
	return len(ph)
}

func (ph PresentHeap) Less(i, j int) bool {
	if ph[i].Value == ph[j].Value {
		return ph[i].Size < ph[j].Size
	}
	return ph[i].Value >= ph[j].Value
}

func (ph PresentHeap) Swap(i, j int) {
	ph[i], ph[j] = ph[j], ph[i]
}

func (ph *PresentHeap) Push(pr interface{}) {
	*ph = append(*ph, pr.(Present))
}

func (ph *PresentHeap) Pop() interface{} {
	old := *ph
	n := len(old)
	pr := old[n-1]
	*ph = old[:n-1]
	return pr
}

func BuildHeap(arr []Present) *PresentHeap {
	res := &PresentHeap{}
	for _, present := range arr {
		heap.Push(res, present)
	}
	return res
}

func GetNCoolestPresents(prArr []Present, n int) ([]Present, error) {
	if n < 0 || n > len(prArr) {
		return nil, fmt.Errorf("n should be in [0; len(arr)], got: %d", n)
	}
	prHeap := BuildHeap(prArr)
	res := make([]Present, n)
	for i := 0; i < n; i++ {
		res[i] = (*prHeap)[i]
	}
	return res, nil
}
