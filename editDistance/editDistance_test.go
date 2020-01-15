package editDistance

import (
	"testing"
	"fmt"
)

func Test1(t *testing.T) {
	res := editDistanceRecursive("dhell", "djhell")
	fmt.Println(res)
	res2 := editDistanceDP("dhell", "djhell")
	fmt.Println(res2)
}