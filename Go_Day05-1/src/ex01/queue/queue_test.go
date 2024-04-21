package queue_test

import (
	"testing"
	"toyTree/queue"
)

func TestEnqueEmpty(t *testing.T) {
	var que queue.Queue
	que.Enque(1)
	expected := []int{1}

	if len(expected) != len(que.Queue) {
		t.Errorf("Wrong len: got %v, wanted %v", len(que.Queue), len(expected))
	} else {
		for i, val := range expected {
			if val != que.Queue[i] {
				t.Errorf("Slices are not equal: got %v, wanted %v", que.Queue, expected)
				break
			}
		}
	}

}

func TestEnqueSeveralElems(t *testing.T) {
	var que queue.Queue
	que.Enque("aboba")
	que.Enque("privet")
	que.Enque("kak dela")
	expected := []string{"aboba", "privet", "kak dela"}

	if len(expected) != len(que.Queue) {
		t.Errorf("Wrong len: got %v, wanted %v", len(que.Queue), len(expected))
	} else {
		for i, val := range expected {
			if val != que.Queue[i] {
				t.Errorf("Slices are not equal: got %v, wanted %v", que.Queue, expected)
				break
			}
		}
	}

}

func TestEmptyDeque(t *testing.T) {
	var que queue.Queue
	expected := 0

	if res := que.Deque(); res != expected {
		t.Errorf("Deque does't work as expected: got %v, wanted %v", res, expected)
	}
}

func TestDeque(t *testing.T) {
	que := queue.Queue{[]interface{}{1, 2, 3, 4}}
	expected := 1

	if res := que.Deque(); res != expected {
		t.Errorf("Deque does't work as expected: got %v, wanted %v", res, expected)
	}
}
