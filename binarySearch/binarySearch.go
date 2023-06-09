package binarysearch

func BinarySearchForInts(list []int, key int) int {
	//Walk to the solution
	for left, right := 0, len(list); left != right; {
		mid := (left + right) / 2

		switch {
		case key < list[mid]:
			right = mid
		case list[mid] == key:
			return mid
		default:
			left = mid + 1
		}
	}
	return -1
}
