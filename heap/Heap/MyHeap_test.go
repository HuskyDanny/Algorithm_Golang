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
	heap.PrintAll()
	t.Log(heap.RemoveTop())
	heap.Append(13)
	t.Log(heap.RemoveTop())
	t.Log(heap.RemoveTop())
	t.Log(heap.RemoveTop())
	heap.PrintAll()
	heap.Append(132)
	heap.Append(1222)
	heap.Append(13332)
	t.Log(heap.RemoveTop())
	heap.PrintAll()
	heap.Append(132222)
	heap.Append(133332)
	heap.Append(1321112)
	heap.Append(13232312)
	heap.PrintAll()
	t.Log(heap.RemoveTop())
	heap.PrintAll()
}

// func TestHeap2(t *testing.T) {
// 	heap := NewMyHeap(10)

// 	heap.Append(5)
// 	heap.PrintAll()
	
// }