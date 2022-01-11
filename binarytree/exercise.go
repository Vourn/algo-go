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

// 二叉树中的最大路径和
type ResultType struct {
    SinglePath int
    MaxPath int
}
func maxPathSum(root *TreeNode) int {
    result := helper(root)
    return result.MaxPath
}
func helper(root *TreeNode) ResultType {
    if root == nil {
	return ResultType {
	    SinglePath: 0
	    MaxPath: -(1 << 31)
	}
    }
    // divide
    left := helper(root.Left)
    right := helper(root.Right)

    // conquer
    result := ResultType{}
    // 单边最大值
    if left.SinglePath > right.SinglePath {
	result.SinglePath = max(left.SinglePath + root.Val, 0)
    } else {
	result.SinglePath = max(right.SinglePath + root.Val, 0)
    }
    // 路径最大值
    maxPath := max(left.MaxPath, right.MaxPath)
    result.MaxPath = max(maxPath, left.SinglePath + right.SinglePath + root.Val)
    return result
}
func max(a, b int) int {
    if a > b {
	return a
    }
    return b
}

// 二叉树的最近公共祖先 (节点在同一子树，返回最先访问到的节点，分布在左右子树，返回父节点)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil {
	return root
    }
    if root.Val == p || root.Val == q {
	return root
    }
    // divide
    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)

    // conquer
    if left != nil && right != nil {
	return root
    }
    if left == nil {
 	return right
    }
    return left
}
