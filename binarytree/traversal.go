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

// 前序非递归
func PreorderTra(root *TreeNode) []int {
    if root == nil {
	return nil
    }
    result := make([]int, 0)
    stack := make([]*TreeNode, 0) // 保存已访问结点，用于原路返回

    for root != nil || len(stack) != 0 {
	for root != nil {
	    // 先保存结果再指向左结点
	    result = append(result, root.Val)
	    stack = append(stack, root)
	    root = root.Left
	}
        // 遍历完左结点，遍历右结点
	node := stack[len(stack) - 1]
        stack = stack[: len(stack) - 1]
	root = root.Right
    }
    return result
}

// 中序非递归

// 后序非递归

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
