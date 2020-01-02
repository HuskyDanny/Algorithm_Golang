package binarySearch



// 13566789

func binarySearchFirstTarget(arr []int, target int) int {

	low := 0
	high := len(arr)

	for low <= high {
		mid := low + (high - low) / 2

		if arr[mid] == target {

			for mid - 1 >= 0 && arr[mid - 1] == target {
				mid -= 1
			}

			return mid

		}else if arr[mid] < target {
			low = mid + 1
		}else {
			high = mid - 1
		}


	}
	return -1
}

func binarySearchLastTarget(arr []int, target int) int {

	low := 0
	high := len(arr)

	for low <= high {
		mid := low + (high - low) / 2

		if arr[mid] == target {

			for mid + 1 < len(arr) && arr[mid + 1] == target {
				mid += 1
			}

			return mid

		}else if arr[mid] < target {
			low = mid + 1
		}else {
			high = mid - 1
		}


	}
	return -1
}

func binarySearchFirstGreater(arr []int, target int) int {
	low := 0
	high := len(arr)

	for low <= high {
		mid := low + (high - low) / 2

		if arr[mid] >= target {

			if mid == 0 || arr[mid-1] < target {
				return mid
			}
			high = mid - 1
		}else {
			low = mid + 1
		}


	}
	return -1
}


func binarySearchLastSmaller(arr []int, target int) int {

	low := 0
	high := len(arr)

	for low <= high {
		mid := low + (high - low) / 2

		if arr[mid] <= target {
			
			if mid == len(arr)-1 || arr[mid + 1] > target {
				return mid
			}

			low = mid + 1

		}else {
			high = mid - 1
		}


	}
	return -1
}