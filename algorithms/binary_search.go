package algorithms

func BinarySearch(items []int, target int) int {
	high := len(items) - 1
	low := 0

	for low <= high {
		guess := (low + high) / 2
		num := items[guess]

		if num == target {
			return guess
		} else if num > target {
			high = guess - 1
		} else {
			low = guess + 1
		}
	}

	return -1
}
