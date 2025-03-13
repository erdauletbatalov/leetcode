package main

// QuickSort выполняет сортировку массива in-place
func QuickSort(arr []int, low, high int) {
	if low < high {
		// Находим индекс опорного элемента после разделения
		pivotIndex := partition(arr, low, high)

		// Рекурсивно сортируем элементы до и после опорного
		QuickSort(arr, low, pivotIndex-1)
		QuickSort(arr, pivotIndex+1, high)
	}
}

// partition разделяет массив на две части относительно опорного элемента
func partition(arr []int, low, high int) int {
	// Выбираем опорный элемент (в данном случае — последний элемент)
	pivot := arr[high]

	// Индекс для элемента, который будет меньше опорного
	i := low - 1

	// Проходим по массиву и перемещаем элементы меньше опорного влево
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i] // Меняем местами
		}
	}

	// Помещаем опорный элемент на правильное место
	arr[i+1], arr[high] = arr[high], arr[i+1]

	// Возвращаем индекс опорного элемента
	return i + 1
}
