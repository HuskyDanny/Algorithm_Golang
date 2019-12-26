package CircularQueue

import "testing"

func TestListQueue(t *testing.T) {
	q := NewCircularQueue(10)
	q.Enqueue(5)

	q.Enqueue(10)

	q.PrintAll()
}