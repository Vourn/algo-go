// leet code exercise
package binarytree
import (
    "math"
)

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

// 给定二叉树，返回其节点值自底向上的层次遍历。（层次遍历+结果翻转）
func levelOrderBottom(root *TreeNode) [][]int {
    result := make([][]int, 0)
    if root == nil {
	return result
    }
    queue := []*TreeNode{root}
    for len(queue) > 0 {
	list := make([]int, 0)
        l := len(queue)
        for i := 0; i < l; i++ {
	    node := queue[0]
	    queue = queue[1:]
	    list = append(list, node.Val)
	    if node.Left != nil {
	    	queue = append(queue, node.Left)
	    }
	    if node.Right != nil {
		queue = append(queue, node.Right)
	    }
	}
    	result = append(result, list)
    }
    for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
	result[i], result[j] = result[j], result[i]
    }
    return result
}

// 给定二叉树，返回其节点值的锯齿形层次遍历，Z字遍历（层次遍历+奇数层节点值翻转）
func zigzagLevelOrder(root *TreeNode) [][]int {
    result := make([][]int, 0)
    if root == nil {
	return result
    }
    queue := []*TreeNode{root}
    toggle := false
    for len(queue) > 0 {
	list := make([]int, 0)
        l := len(queue)
	for i := 0; i < l; i++ {
	    // 节点出队
	    node := queue[0]
	    queue = queue[1:]
	    list = append(list, node.Val)
	    if node.Left != nil {                
    		queue = append(queue, node.Left) 
	    }                                    
	    if node.Right != nil {               
    		queue = append(queue, node.Right)
	    }                                    
	}
        if toggle {
	     for j := 0; j < len(list)/2; j++ {
		list[i], list[len(list)-1-i] = list[len(list)-1-i], list[i]	
	     }
	}
        result = append(result, list)
        toggle = !toggle
    }
    return result
}

// 判断是否为有效二叉搜索树（左节点<根节点<右节点）
func isValidBST(root *TreeNode) bool {
    return helperBST(root, math.MinInt64, math.MaxInt64)
}

func helperBST(root *TreeNode, lower, upper int) bool {
    if root == nil {
	return true
    }
    if root.Val <= lower || root.Val >= upper {
	return false
    }
    return helperBST(root.Left, lower, root.Val) && helperBST(root.Right, roo.Val, upper)
}

// 二叉搜索树插入操作
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
    	root = &TreeNode{Val: val}
	return root
    }
    if val < root.Val {
	root.Left = insertIntoBST(root.Left, val)
    } else {
	root.Right = insertIntoBST(root.Right, val)
    }
    return root
}

// 二叉树删除操作
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
	return root
    }
    if root.Val == key {
  	if root.Left == nil {
 	    return root.Right
	} else if root.Right == nil {
	    return root.Left
	} else if {
	    cur := root.Right
	    // 找到左分支最后一个节点
	    for cur.Left != nil {
		cur = cur.Left
	    }
	    return root.Right
	}
    }
    if root.Val > key {
	root.Left = deleteNode(root.Left, key)
    } else {
	root.Right = deleteNode(root.Right, key)
    }
    return root
}
