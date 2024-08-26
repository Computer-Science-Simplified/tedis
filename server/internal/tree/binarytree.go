package tree

import (
	"fmt"
	"github.com/Computer-Science-Simplified/tedis/server/internal/enum"
)

type BinaryTree struct {
	Key  string
	Root *BinaryTreeNode
}

type BinaryTreeNode struct {
	Value int64
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func (tree *BinaryTree) GetKey() string {
	return tree.Key
}

func (tree *BinaryTree) GetType() string {
	return enum.BinaryTree
}

func (tree *BinaryTree) Add(value int64) {
	tree.add(value, tree.Root)
}

func (tree *BinaryTree) Exists(value int64) bool {
	return tree.exists(value, tree.Root)
}

func (tree *BinaryTree) Remove(value int64) {
	fmt.Println("remove")
}

func (tree *BinaryTree) GetAll() []int64 {
	queue := make([]*BinaryTreeNode, 0)

	if tree.Root != nil {
		queue = append(queue, tree.Root)
	}

	values := make([]int64, 0)

	for len(queue) > 0 {
		current := queue[0]

		queue = queue[1:]

		values = append(values, current.Value)

		if current.Left != nil {
			queue = append(queue, current.Left)
		}

		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return values
}

// -------- Private functions --------

func (tree *BinaryTree) add(value int64, node *BinaryTreeNode) {
	if node == nil {
		newNode := BinaryTreeNode{Value: value}
		tree.Root = &newNode
		return
	}

	queue := make([]*BinaryTreeNode, 0)

	if tree.Root != nil {
		queue = append(queue, tree.Root)
	}

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		newNode := &BinaryTreeNode{Value: value}

		if node.Left == nil {
			node.Left = newNode
			return
		}

		if node.Right == nil {
			node.Right = newNode
			return
		}

		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}
}

func (tree *BinaryTree) exists(value int64, node *BinaryTreeNode) bool {
	if node == nil {
		return false
	}

	if node.Value == value {
		return true
	}

	existsLeft := tree.exists(value, node.Left)

	if existsLeft {
		return true
	}

	existsRight := tree.exists(value, node.Right)

	return existsRight
}
