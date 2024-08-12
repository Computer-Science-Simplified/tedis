package trees

import (
	"mmartinjoo/trees/commands"

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

func (tree *BST) Insert(value int64, node *BSTNode, shouldReport bool) *BSTNode {
	if shouldReport {
		event.MustFire("write_command_executed", event.M{
			"command": commands.BTADD,
			"key":     tree.Key,
			"args":    []int64{value},
		})
	}

	if tree.Root == nil {
		newNode := BSTNode{Value: value}

		tree.Root = &newNode

		return &newNode
	}

	return tree.insert(value, tree.Root)
}

func (tree *BST) Exists(value int64) bool {
	return tree.exists(value, tree.Root)
}

func (tree *BST) Remove(value int64) {
	event.MustFire("write_command_executed", event.M{
		"command": commands.BTADD,
		"key":     tree.Key,
		"args":    []int64{value},
	})

	tree.remove(value, tree.Root, nil)
}

func (tree *BST) ToArray() []int64 {
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

func (tree *BST) insert(value int64, node *BSTNode) *BSTNode {
	if node == nil {
		newNode := BSTNode{Value: value}

		return &newNode
	}

	if value < node.Value {
		node.Left = tree.insert(value, node.Left)
	} else {
		node.Right = tree.insert(value, node.Right)
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

func (tree *BST) remove(value int64, node *BSTNode, parent *BSTNode) {
	if node == nil {
		return
	}

	if node.Value == value {
		if parent == tree.Root || parent == nil {
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
