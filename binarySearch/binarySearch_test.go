package binarySearch

import (
	"testing"
	"fmt"
)

func TestBSFirstTarget(t *testing.T) {

	arr1  := []int {1,2,2,3,3,3,3,3,4,4,5,6}
	
	fmt.Println(binarySearchFirstTarget(arr1, 3))
}


func TestBSLastTarget(t *testing.T) {

	arr1  := []int {1,2,2,3,3,3,3,3,4,4,5,6}
	
	fmt.Println(binarySearchLastTarget(arr1, 3))
}

func TestBSFirstGreater(t *testing.T) {

	arr1  := []int {3,4,6,7,10}
	
	fmt.Println(binarySearchFirstGreater(arr1, 8))
}

func TestBSLastSmaller(t *testing.T) {
	arr1  := []int {2,2,2,2,2,2,2,2,2}
	
	fmt.Println(binarySearchLastSmaller(arr1, 3))
}