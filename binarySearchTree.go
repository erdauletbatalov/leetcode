package main

import (
	"fmt"
)

// Узел BST
type BSTNode struct {
	Value int
	Left  *BSTNode
	Right *BSTNode
}

// Вставка элемента в BST
func Insert(root *BSTNode, value int) *BSTNode {
	// Если дерево пустое, создаем новый узел
	if root == nil {
		return &BSTNode{Value: value, Left: nil, Right: nil}
	}

	// Если значение меньше текущего узла, вставляем в левое поддерево
	if value < root.Value {
		root.Left = Insert(root.Left, value)
	} else if value > root.Value {
		// Если значение больше текущего узла, вставляем в правое поддерево
		root.Right = Insert(root.Right, value)
	}

	// Возвращаем неизмененный указатель root
	return root
}

// Обход дерева in-order (для проверки)
func InOrderTraversal(root *BSTNode) {
	if root != nil {
		InOrderTraversal(root.Left)
		fmt.Printf("%d ", root.Value)
		InOrderTraversal(root.Right)
	}
}

// Удаление элемента из BST
func Remove(root *BSTNode, value int) *BSTNode {
	// Если дерево пустое, возвращаем nil
	if root == nil {
		return nil
	}

	// Ищем узел для удаления
	if value < root.Value {
		root.Left = Remove(root.Left, value)
	} else if value > root.Value {
		root.Right = Remove(root.Right, value)
	} else {
		// Узел найден

		// Случай 1: У узла нет дочерних элементов или только один дочерний элемент
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		// Случай 3: У узла два дочерних элемента
		// Находим минимальный элемент в правом поддереве
		minNode := FindMin(root.Right)
		// Копируем значение минимального элемента в текущий узел
		root.Value = minNode.Value
		// Удаляем минимальный элемент из правого поддерева
		root.Right = Remove(root.Right, minNode.Value)
	}

	return root
}

// Поиск минимального элемента в BST
func FindMin(node *BSTNode) *BSTNode {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}
