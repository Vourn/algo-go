// leet code exercise
package binarytree

// 一棵二叉树的最大深度
func MaxDepth(root *TreeNode) int {
    if root == nil {
	return 0
    }
    // 左右子树分别递归计算深度
    left := MaxDepth(root.Left)
    right := MaxDepth(root.Right)
    // 合并左右子树返回数
    if left > right {
 	return left + 1
    }
    return right + 1
}

// 判断是否为高度平衡二叉树
func isBalanced(root *TreeNode) bool {
    if maxDepth(root) == -1 {
	return false
    }
    return true
}
func maxDepth(root *TreeNode) int {
    if root == nil {
	return 0
    }
    left := maxDepth(root.Left)
    right := maxDepth(root.Right)

    if left == -1 || right == -1 || left - right > 1 || right - left > 1 {
	return -1
    }
    if left > right {
	return left + 1
    }
    return right + 1
}
