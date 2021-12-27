package queue_test

import (
	"testing"

	"github.com/st3fan/queue"
)

func Test_SynchronizedQueueImplementsQueue(t *testing.T) {
	var _ queue.QueueInterface[int] = &queue.SynchronizedQueue[int]{}
}
