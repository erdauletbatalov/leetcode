package main

import "fmt"

func main() {
	var root *BSTNode

	// Вставляем элементы в BST
	root = Insert(root, 10)
	root = Insert(root, 5)
	root = Insert(root, 15)
	root = Insert(root, 3)
	root = Insert(root, 7)
	root = Insert(root, 12)
	root = Insert(root, 18)

	// Обход дерева in-order до удаления
	fmt.Println("In-order traversal before removal:")
	InOrderTraversal(root)
	fmt.Println()

	// Удаляем элемент 5
	root = Remove(root, 5)

	// Обход дерева in-order после удаления
	fmt.Println("In-order traversal after removal:")
	InOrderTraversal(root)
	fmt.Println()
}
