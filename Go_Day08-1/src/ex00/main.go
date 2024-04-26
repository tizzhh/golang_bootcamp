package main

import (
	"fmt"
	"unsafe"
)

func getElement(arr []int, idx int) (int, error) {
	if len(arr) == 0 {
		return 0, fmt.Errorf("empty slice")
	}
	if idx < 0 || idx >= len(arr) {
		return 0, fmt.Errorf("idx out of bounds: %d", idx)
	}

	return *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&arr[0])) + unsafe.Sizeof(int(0))*uintptr(idx))), nil
}

func main() {
	fmt.Println(getElement([]int{1, 2, 3}, 1))
}
