package main

import (
	"sort"
)

// BucketSort реализует сортировку методом корзин (Bucket Sort)
func BucketSort(arr []float64, bucketSize int) []float64 {
	if len(arr) == 0 {
		return arr
	}

	// 1️⃣ Найти минимум и максимум массива
	minVal, maxVal := arr[0], arr[0]
	for _, num := range arr {
		if num < minVal {
			minVal = num
		} else if num > maxVal {
			maxVal = num
		}
	}

	// 2️⃣ Определить количество корзин
	bucketCount := (int(maxVal-minVal)/bucketSize + 1)
	buckets := make([][]float64, bucketCount)

	// 3️⃣ Распределить элементы по корзинам
	for _, num := range arr {
		idx := int((num - minVal) / float64(bucketSize))
		buckets[idx] = append(buckets[idx], num)
	}

	// 4️⃣ Отсортировать каждую корзину (используем встроенную sort.Float64s)
	sortedArr := []float64{}
	for _, bucket := range buckets {
		sort.Float64s(bucket)
		sortedArr = append(sortedArr, bucket...)
	}

	return sortedArr
}
