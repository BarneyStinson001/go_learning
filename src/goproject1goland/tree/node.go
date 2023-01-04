package tree

import "fmt"

type point struct {
	i,j int
}

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func CreateNode(value int)*TreeNode {
	return &TreeNode{Value: value} //局部变量也能返回实用
}

func (node TreeNode) Print()  {
	fmt.Println(node.Value)
}

func (node TreeNode)SetValue(value int)  {
	node.Value=value
}
func (node *TreeNode)SetValue2(value int)  {
	node.Value=value
}


