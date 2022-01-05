package binarytree

// 前序递归
func PreorderTraversal(root *TreeNode) {
    if root == nil {
	return
    }
    // 根左右
    fmt.Println(root.Val)
    PreorderTraversal(root.Left)
    PreorderTraversal(root.Right)
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
