package queue

type Queue[T any] interface {
	Push(input T)
	Pop() (T, bool)
	Length() int
}

type node[T any] struct {
	value T
	prev  *node[T]
}

type queue[T any] struct {
	head, tail *node[T]
	length     int
}

func New[T any]() Queue[T] {
	return &queue[T]{}
}

func (q *queue[T]) Push(input T) {
	newV := &node[T]{value: input}
	if q.head == nil {
		q.head = newV
		q.tail = newV
	} else {
		q.tail.prev = newV
		q.tail = newV
	}
	q.length++
}

func (q *queue[T]) Pop() (T, bool) {
	if q.head == nil {
		var empty T
		q.length = 0
		return empty, false
	}
	out := q.head.value
	q.head = q.head.prev
	q.length--
	return out, true
}

func (q *queue[T]) Length() int {
	return q.length
}
