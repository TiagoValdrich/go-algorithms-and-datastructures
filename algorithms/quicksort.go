package algorithms

func Quicksort(nums []int) []int {
	size := len(nums)

	if size < 2 {
		return nums
	}

	pivot := nums[0]

	var (
		smaller []int
		bigger  []int
	)

	for i := 1; i < size; i++ {
		if nums[i] > pivot {
			bigger = append(bigger, nums[i])
		} else {
			smaller = append(smaller, nums[i])
		}
	}

	var sortedList []int

	smallerQs := Quicksort(smaller)
	biggerQs := Quicksort(bigger)
	sortedList = append(sortedList, smallerQs...)
	sortedList = append(sortedList, pivot)
	sortedList = append(sortedList, biggerQs...)

	return sortedList
}
