package main

// Sort an Array

// Topics:
//
// Array
// Divide and Conquer
// Sorting
// Heap (Priority Queue)
// Merge Sort
// Bucket Sort
// Radix Sort
// Counting Sort

func sortArray(nums []int) []int {
	// Merge Sort
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2

	left := sortArray(nums[:mid])
	right := sortArray(nums[mid:])

	return merge912(left, right)

	// // Insertion Sort
	// for i := 0; i < len(nums); i++ {
	// 	j := i - 1
	// 	for j >= 0 && nums[j] > nums[j+1] {
	// 		nums[j], nums[j+1] = nums[j+1], nums[j]
	// 		j--
	// 	}
	// }
	// return nums
}

func merge912(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
