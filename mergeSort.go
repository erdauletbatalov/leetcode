package main

// MergeSort выполняет сортировку слиянием
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Разделяем массив пополам
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	// Объединяем отсортированные половины
	return merge(left, right)
}

// merge сливает два отсортированных массива
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// Пока есть элементы в обоих массивах
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Добавляем оставшиеся элементы
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
