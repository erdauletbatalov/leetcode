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
	for i := 0; i < len(nums); i++ {
		j := i - 1
		for j >= 0 && nums[j] > nums[j+1] {
			nums[j], nums[j+1] = nums[j+1], nums[j]
			j--
		}
	}
	return nums
}
