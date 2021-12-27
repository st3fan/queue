package queue

import "sync"

type SynchronizedQueue[T any] struct {
	mutex sync.Mutex
	queue QueueInterface[T]
}

func (q *SynchronizedQueue[T]) Add(e T) error {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return q.queue.Add(e)
}

func (q *SynchronizedQueue[T]) Remove() (T, bool) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return q.queue.Remove()
}

func (q *SynchronizedQueue[T]) Peek() (T, bool) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return q.queue.Peek()
}

func (q *SynchronizedQueue[T]) Length() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return q.queue.Length()
}

func (q *SynchronizedQueue[T]) Clear() {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.queue.Clear()
}
