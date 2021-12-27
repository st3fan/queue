package queue

import "github.com/st3fan/deque"

type QueueInterface[T any] interface {
	Add(e T) error
	Remove() (T, bool)
	Peek() (T, bool)
	Length() int
	Clear()
}

//

type Queue[T any] struct {
	deque deque.Deque[T]
}

func (q *Queue[T]) Add(e T) error {
	return q.deque.AddLast(e)
}

func (q *Queue[T]) Remove() (T, bool) {
	return q.deque.RemoveFirst()
}

func (q *Queue[T]) Peek() (T, bool) {
	return q.deque.PeekFirst()
}

func (q *Queue[T]) Length() int {
	return q.deque.Length()
}

func (q *Queue[T]) Clear() {
	q.deque.Clear()
}
