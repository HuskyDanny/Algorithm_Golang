package MyHeap

import "testing"

func TestHeap(t *testing.T) {
	heap := NewMyHeap(10)

	heap.Append(5)
	heap.Append(10)
	heap.Append(9)
	heap.Append(8)
	heap.Append(11)
	heap.PrintAll()
	t.Log(heap.RemoveTop())
	t.Log(heap.RemoveTop())
	heap.Append(13)
	t.Log(heap.RemoveTop())
	t.Log(heap.RemoveTop())
	t.Log(heap.RemoveTop())
	heap.Append(132)
	heap.PrintAll()
}