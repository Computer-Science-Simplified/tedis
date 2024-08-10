package trees

import (
	"fmt"

	"github.com/gookit/event"
)

type BinaryTree struct {
	Key string
	Root *BinaryTreeNode
}

type BinaryTreeNode struct {
    Value int64
    Left  *BinaryTreeNode
    Right *BinaryTreeNode
}

func (tree *BinaryTree) Insert(value int64, node *BinaryTreeNode) *BinaryTreeNode {
	event.MustFire("write_command_executed", event.M{
		"command": "BTADD", 
		"key": tree.Key, 
		"args": []int64{value},
	})

	if tree.Root == nil {
		newNode := BinaryTreeNode{Value: value}

		tree.Root = &newNode

		return &newNode
	}

	return tree.insert(value, tree.Root)
}

func (tree *BinaryTree) Exists(value int64) bool {
	return tree.exists(value, tree.Root)
}

func (tree *BinaryTree) Remove(value int64) {
	tree.remove(value, tree.Root, nil)
}

func (tree *BinaryTree) ToArray() []int64 {
	queue := make([]*BinaryTreeNode, 0)

	queue = append(queue, tree.Root)

	values := make([]int64, 0)

	for len(queue) != 0 {
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

func (tree *BinaryTree) insert(value int64, node *BinaryTreeNode) *BinaryTreeNode {
	if node == nil {
		newNode := BinaryTreeNode{Value: value}

		return &newNode
	}

	if value < node.Value {
		node.Left = tree.insert(value, node.Left)
	} else {
		node.Right = tree.insert(value, node.Right)
	}

	return node
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

func (tree *BinaryTree) remove(value int64, node *BinaryTreeNode, parent *BinaryTreeNode) {
	if node == nil {
		return
	}

	if node.Value == value {
		if parent == tree.Root || parent == nil {
			fmt.Println(value)
			tree.Root = nil

			return
		}

		if value < parent.Value {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
	}

	tree.remove(value, node, node.Left)
	tree.remove(value, node, node.Right)
}