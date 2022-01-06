package main

import (
    "fmt"
    "github.com/Vourn/algo-go/binarytree"
)
func main() {
    nodeH := &binarytree.TreeNode {
	Val: 8,
	Left: nil,
	Right: nil,	
    }
    nodeG := &binarytree.TreeNode {
        Val: 7,                   
    	Left: nil,                
    	Right: nil,               
    }
    nodeF := &binarytree.TreeNode {
    	Val: 6,                   
    	Left: nil,                
    	Right: nil,               
    }
    nodeE := &binarytree.TreeNode {
    	Val: 5,                   
    	Left: nil,                
    	Right: nil,               
    }
    nodeD := &binarytree.TreeNode {
    	Val: 4,                   
    	Left: nil,                
    	Right: nodeH,               
    }
    nodeC := &binarytree.TreeNode{
    	Val: 3,                   
    	Left: nodeF,                
    	Right: nodeG,               
    }
    nodeB := &binarytree.TreeNode {
	Val: 2,
	Left: nodeD,
	Right: nodeE,
    }
    nodeA := &binarytree.TreeNode {
	Val: 1,
	Left: nodeB,
	Right: nodeC,
    }
    preResult := binarytree.PreorderTra(nodeA)
    inResult := binarytree.InorderTra(nodeA)
    backResult := binarytree.BackorderTra(nodeA)
    fmt.Println(preResult)
    fmt.Println(inResult)
    fmt.Println(backResult)
}
