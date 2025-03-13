package main

type TreeNode struct {
	val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTreeDFS(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = invertTreeDFS(root.Left)
	root.Right = invertTreeDFS(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

func invertTreeBFS(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return root
}

func maxDepthDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepthDFS(root.Left), maxDepthDFS(root.Right))
}

func maxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	depth := 0
	for len(queue) > 0 {
		depth++
		for i := len(queue); i > 0; i-- {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return depth
}
