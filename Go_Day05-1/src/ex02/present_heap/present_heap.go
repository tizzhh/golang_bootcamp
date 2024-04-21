package presentHeap

import "container/heap"

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
	return ph[i].Value > ph[j].Value
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

func BuildHeapByPush(arr []Present) *PresentHeap {
	res := &PresentHeap{}
	for _, elem := range arr {
		heap.Push(res, elem)
	}
	return res
}

func GetNCoolestPresents(prArr []Present, n int) []Present {
	prHeap := BuildHeapByPush(prArr)
	var res []Present
	for _, present := range *prHeap {
		if present.Value == n {
			res = append(res, present)
		}
	}
	return res
}
