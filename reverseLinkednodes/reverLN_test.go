package reverseLinkednodes

import ("testing"
		"fmt"
)


func TestReverseLN(t *testing.T) {
	t1 := LinkedNode{nil, 1}
	t2 := LinkedNode{nil, 2}
	t3 := LinkedNode{nil, 3}
	t4 := LinkedNode{nil, 4}
	t5 := LinkedNode{nil, 5}
	t1.next = &t2
	t2.next = &t3
	t3.next = &t4
	t4.next = &t5


	
	res := ReverseRecursive(&t1)
	fmt.Println("  ")
	res.PrintAll()
	
}