package binarytree

import (
    "fmt"
)

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
	root = node.Right
    }
    return result
}

// 中序非递归
func InorderTra(root *TreeNode) []int {
    if root == nil {
	return nil
    }
    result := make([]int, 0)
    stack := make([]*TreeNode, 0)

    for root != nil || len(stack) != 0 {
	for root != nil {
	    stack = append(stack, root)
	    root = root.Left // 一直往左
	}
        // 左结点遍历完，原路依次遍历根右节点
	node := stack[len(stack) - 1]
	stack = stack[: len(stack) -1]
        result = append(result, node.Val)
	root = node.Right
    }
    return result
}

// 后序非递归
func BackorderTra(root *TreeNode) []int {
    // lastVisit标识右子结点是否已弹出
    if root == nil {
	return nil
    }
    result := make([]int, 0)
    stack := make([]*TreeNode, 0)
    var lastVisit *TreeNode

    for root != nil || len(stack) != 0 {
	for root != nil {
	    stack = append(stack, root)
	    root = root.Left // 一直往左
	}
	// 原路访问右根结点
	node := stack[len(stack) -1]
	if node.Right == nil || node.Right == lastVisit {
	    stack = stack[: len(stack) -1]
	    result = append(result, node.Val)	    
	    // 标记弹出
	    lastVisit = node
	} else {
	    root = node.Right
	}
    }
    return result
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
