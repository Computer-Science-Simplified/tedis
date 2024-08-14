package tree

import (
	"mmartinjoo/trees/internal/enum"

	"github.com/gookit/event"
)

type BST struct {
	Key  string
	Root *BSTNode
}

type BSTNode struct {
	Value int64
	Left  *BSTNode
	Right *BSTNode
}

func (tree *BST) GetKey() string {
	return tree.Key
}

func (tree *BST) GetType() string {
	return enum.BinarySearchTree
}

func (tree *BST) Add(value int64, shouldReport bool) {
	if shouldReport {
		event.MustFire("write_command_executed", event.M{
			"command": enum.BSTADD,
			"key":     tree.Key,
			"args":    []int64{value},
		})
	}

	tree.Root = tree.add(value, tree.Root)
}

func (tree *BST) Exists(value int64) bool {
	return tree.exists(value, tree.Root)
}

func (tree *BST) Remove(value int64, shouldReport bool) {
	if shouldReport {
		event.MustFire("write_command_executed", event.M{
			"command": enum.BSTREM,
			"key":     tree.Key,
			"args":    []int64{value},
		})
	}

	tree.remove(value, tree.Root)
}

func (tree *BST) GetAll() []int64 {
	queue := make([]*BSTNode, 0)

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

func (tree *BST) add(value int64, node *BSTNode) *BSTNode {
	if node == nil {
		newNode := BSTNode{Value: value}

		return &newNode
	}

	if value < node.Value {
		node.Left = tree.add(value, node.Left)
	} else {
		node.Right = tree.add(value, node.Right)
	}

	return node
}

func (tree *BST) exists(value int64, node *BSTNode) bool {
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

func (tree *BST) remove(value int64, node *BSTNode) *BSTNode {
	if node == nil {
		return nil
	}

	if node.Value == value {
		// No child or only one child
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		// Two children
		smallestNode := tree.findSmallestNode(node.Right)
		node.Value = smallestNode.Value
		node.Right = tree.remove(smallestNode.Value, node.Right)

	} else if value < node.Value {
		node.Left = tree.remove(value, node.Left)
	} else {
		node.Right = tree.remove(value, node.Right)
	}

	return node
}

func (tree *BST) findSmallestNode(node *BSTNode) *BSTNode {
	current := node

	for current.Left != nil {
		current = current.Left
	}

	return current
}
