package main

import "fmt"
import "goproject1goland/tree"
import "goproject1goland/queue"
type myTreeNode struct {
	//node *tree.TreeNode  //用指针避免拷贝一份
	 *tree.TreeNode  //语法糖省略 node

}

func (myNode *myTreeNode) postOrder()  {
	//if myNode == nil || myNode.node== nil{
	if myNode == nil || myNode.TreeNode== nil{
			return
	}
	//left := myTreeNode{myNode.node.Left}
	left := myTreeNode{myNode.Left}//还省了.TreeNode
	//right := myTreeNode{myNode.node.Right}
	right := myTreeNode{myNode.Right}//还省了.TreeNode

	left.postOrder()
	right.postOrder()
	myNode.Print()//还省了.TreeNode
}

func main() {
	//var root tree.TreeNode//treenode 小写表示私有，要改成大写
	//root=tree.TreeNode{Value: 3}
	root:=myTreeNode{&tree.TreeNode{3,nil,nil}}
	root.Left=&tree.TreeNode{} //少这个的话，后面关于root.left的都会报错失败 。
	root.Right= &tree.TreeNode{5,nil,nil}
	root.Right.Left=new(tree.TreeNode)
	root.Left.Right=tree.CreateNode(2)
	root.Right.Left.SetValue2(4)
	root.Traveling()
    //postorder := myTreeNode{&root}
    root.postOrder()
	nodes := []tree.TreeNode{
		{Value: 3},
		{},
		{6,nil,nil},
	}
	fmt.Println(nodes)
	fmt.Println("==============1")



	root.Left.Right=tree.CreateNode(2)
	fmt.Println("==============2")

	root.Print()
	root.Left.SetValue(4)//值传递不改变原有值
	root.Left.Print()
	root.Left.SetValue2(4)//定义的指针可以修改到原有值
	root.Left.Print()
	pRoot :=&root
	pRoot.Print()
	pRoot.SetValue2(100)
	pRoot.Print()

	fmt.Println("==============0")
	q:=queue.Queue{}
	fmt.Println(q.ISEmpty())
	q.Push(1)
	fmt.Println(q.ISEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.ISEmpty())



}