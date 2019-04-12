package main

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	Value     int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

// NewTree 新建树
func NewTree(values ...int) (*TreeNode, error) {
	if len(values) <= 0 {
		return nil, errors.New("length error")
	}
	rootNode := &TreeNode{
		Value: values[0],
	}
	for i, lenght := 1, len(values); i < lenght; i++ {
		rootNode.setChildren(values[i])
	}
	return rootNode, nil
}

func (node *TreeNode) setChildren(value int) {
	if value <= node.Value {
		node.setLeft(value)
	} else {
		node.setRight(value)
	}
}

func (node *TreeNode) setLeft(value int) {
	if node.LeftNode != nil {
		node.LeftNode.setChildren(value)
	} else {
		node.LeftNode = &TreeNode{
			Value: value,
		}
	}
}

func (node *TreeNode) setRight(value int) {
	if node.RightNode != nil {
		node.RightNode.setChildren(value)
	} else {
		node.RightNode = &TreeNode{
			Value: value,
		}
	}
}

// FrontForeach 先序遍历
func (node *TreeNode) FrontForeach() {
	if node == nil {
		return
	}
	fmt.Printf("%v ", node.Value)
	node.LeftNode.FrontForeach()
	node.RightNode.FrontForeach()
}

// MidForeach 中序遍历
func (node *TreeNode) MidForeach() {
	if node == nil {
		return
	}
	node.LeftNode.MidForeach()
	fmt.Printf("%v ", node.Value)
	node.RightNode.MidForeach()
}

// AfterForeach 后续遍历
func (node *TreeNode) AfterForeach() {
	if node == nil {
		return
	}
	node.LeftNode.AfterForeach()
	node.RightNode.AfterForeach()
	fmt.Printf("%v ", node.Value)
}

// GetHeight 获取深度
func (node *TreeNode) GetHeight() int {
	if node == nil {
		return 0
	}
	left := node.LeftNode.GetHeight()
	right := node.RightNode.GetHeight()

	if left < right {
		return right + 1
	} else {
		return left + 1
	}
}

// GetMax 获取最大值
func (node *TreeNode) GetMax() (int, error) {
	if node == nil {
		return 0, errors.New("none")
	}
	max := node.Value
	left, err := node.LeftNode.GetMax()
	if err == nil && max < left {
		max = left
	}
	right, err := node.RightNode.GetMax()
	if err == nil && max < right {
		max = right
	}
	return max, nil
}

// GetMin 获取最大值
func (node *TreeNode) GetMin() (int, error) {
	if node == nil {
		return 0, errors.New("none")
	}
	min := node.Value
	left, err := node.LeftNode.GetMax()
	if err == nil && min > left {
		min = left
	}
	right, err := node.RightNode.GetMax()
	if err == nil && min > right {
		min = right
	}
	return min, nil
}

func main() {
	arr := []int{10, 9, 20, 35, 15}
	tree, err := NewTree(arr...)
	if err != nil {
		fmt.Println(err)
		return
	}
	tree.FrontForeach()
	fmt.Println()
	tree.MidForeach()
	fmt.Println()
	tree.AfterForeach()
	fmt.Println()
	fmt.Print("当前树深度: ")
	fmt.Println(tree.GetHeight())
	fmt.Print("当前树最大值: ")
	fmt.Println(tree.GetMax())
	fmt.Print("当前树最小值: ")
	fmt.Println(tree.GetMin())
}
