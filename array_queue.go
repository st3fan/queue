package queue

type ArrayQueue[T any] struct {
	data []T
}

func (q *ArrayQueue[T]) Add(e T) error {
	return nil
}

func (q *ArrayQueue[T]) Remove() (T, bool) {
	var empty T
	return empty, false
}

func (q *ArrayQueue[T]) Peek() (T, bool) {
	var empty T
	return empty, false
}

func (q *ArrayQueue[T]) Length() int {
	return 0
}

func (q *ArrayQueue[T]) Clear() {
}
