package ds

//TreeNode 二叉树节点
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

//Node 二叉树节点
type Node struct {
    Val int
    Left *Node
    Right *Node
    //指向同层的下一个节点
    Next *Node
}