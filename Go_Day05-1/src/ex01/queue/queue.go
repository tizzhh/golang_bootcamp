package queue

type Queue struct {
	Queue []interface{}
}

func (q *Queue) Enque(elem interface{}) {
	q.Queue = append(q.Queue, elem)
}

func (q *Queue) Deque() interface{} {
	if len(q.Queue) == 0 {
		return 0
	}
	elem := q.Queue[0]
	if len(q.Queue) == 1 {
		q.Queue = nil
	} else {
		q.Queue = q.Queue[1:]
	}

	return elem
}
