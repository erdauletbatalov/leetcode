package main

// MinHeap представляет минимальную кучу
type MinHeap struct {
	heap []int
}

// Insert добавляет элемент в кучу
func (h *MinHeap) Insert(value int) {
	h.heap = append(h.heap, value) // Добавляем элемент в конец массива
	h.heapifyUp(len(h.heap) - 1)   // Восстанавливаем свойство кучи
}

// ExtractMin извлекает и возвращает минимальный элемент
func (h *MinHeap) ExtractMin() int {
	if len(h.heap) == 0 {
		panic("Heap is empty")
	}

	min := h.heap[0] // Минимальный элемент — это корень
	lastIndex := len(h.heap) - 1
	h.heap[0] = h.heap[lastIndex] // Перемещаем последний элемент в корень
	h.heap = h.heap[:lastIndex]   // Удаляем последний элемент
	h.heapifyDown(0)              // Восстанавливаем свойство кучи
	return min
}

// heapifyUp восстанавливает свойство кучи снизу вверх
func (h *MinHeap) heapifyUp(index int) {
	for h.heap[parent(index)] > h.heap[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// heapifyDown восстанавливает свойство кучи сверху вниз
func (h *MinHeap) heapifyDown(index int) {
	lastIndex := len(h.heap) - 1
	left, right := leftChild(index), rightChild(index)
	childToCompare := 0

	// Цикл выполняется, пока у узла есть хотя бы один дочерний элемент
	for left <= lastIndex {
		if left == lastIndex { // Если есть только левый дочерний элемент
			childToCompare = left
		} else if h.heap[left] < h.heap[right] { // Выбираем меньший из двух дочерних элементов
			childToCompare = left
		} else {
			childToCompare = right
		}

		// Если текущий узел больше меньшего дочернего элемента, меняем их местами
		if h.heap[index] > h.heap[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			left, right = leftChild(index), rightChild(index)
		} else {
			return
		}
	}
}

// Вспомогательные функции для работы с индексами
func parent(index int) int {
	return (index - 1) / 2
}

func leftChild(index int) int {
	return 2*index + 1
}

func rightChild(index int) int {
	return 2*index + 2
}

// swap меняет местами два элемента в массиве
func (h *MinHeap) swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

// func main() {
// 	heap := &MinHeap{}

// 	// Вставляем элементы в кучу
// 	heap.Insert(10)
// 	heap.Insert(20)
// 	heap.Insert(30)
// 	heap.Insert(5)
// 	heap.Insert(15)

// 	fmt.Println("Heap after insertions:", heap.heap)

// 	// Извлекаем минимальные элементы
// 	fmt.Println("Extracted min:", heap.ExtractMin())
// 	fmt.Println("Heap after extraction:", heap.heap)

// 	fmt.Println("Extracted min:", heap.ExtractMin())
// 	fmt.Println("Heap after extraction:", heap.heap)
// }
