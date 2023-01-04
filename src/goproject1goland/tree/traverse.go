package tree


func (node *TreeNode) Traveling()  {
	if node ==nil{
		return
	}
	node.Left.Traveling()
	node.Print()
	node.Right.Traveling()
}

