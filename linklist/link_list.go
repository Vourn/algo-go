// 链表练习
package linklist

// 删除升序链表中的重复元素,使每个元素只出现一次
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
	return nil
    }
    cur := head
    for cur.Next != nil {
	if cur.Val == cur.Next.Val {
	    cur.Next = cur.Next.Next
        } else {
	    cur = cur.Next
        }
    }
    return head
}

// 删除升序链表中的重复元素，只保留原链表中没有重复出现的数字
func deleDuplicates(head *ListNode) *ListNode {
    if head == nil {
	return head
    }
    // head 节点可能会被删除，用dummy node辅助
    dummy := &ListNode{0, head}
    cur := dummy
    for cur.Next != nil && cur.Next.Next != nil {
	if cur.Next.Val == cur.Next.Next.Val {
	    // 查找重复值
	    rm := cur.Next.Val
	    for cur.Next != nil && cur.Next.Val == rm {
		cur.Next = cur.Next.Next
	    }
	} else {
	    cur = cur.Next
	}
    }
    return dummy.Next
}

// 反转链表
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    for head != nil {
	temp := head.Next
        head.Next = prev
        prev = head
	head = temp
    }
    return prev
}

// 反转从位置left到right的链表。请使用一趟扫描完成反转
func reverseBetween(head *ListNode, left, right int) *ListNode {
    dummy := &ListNode{Val: 0}
    dummy.Next = head
    prev := dummy
    // 遍历到left位置
    for i := 0; i < left -1; i++ {
	prev = prev.Next
    } 
    cur := prev.Next
    // left到right区间节点反转
    for i := 0; i < right - left; i++ {
	next := cur.Next
        cur.Next = next.Next
        next.Next = prev.Next
	prev.Next = next
    }
    return dummy.Next
}

// 合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{Val: 0}
    head := dummy
    for list1 != nil && list2 != nil {
	if list1.Val < head.Val {
	    head.Next = list1
	    list1 = list1.Next
	} else {
	    head.Next = list2
	    list2 = list2.Next
        }
	head = head.Next
    }
    // 合并剩余部分
    if list1 != nil {
	head.Next = list1
    } else {
 	head.Next = list2
    }
    return dummy.Next
}

// 分隔链表，使小于x的节点都在大于等于x节点之前（新链表存储大于等于的节点，最后连接两个链表）
func partition(head *ListNode, x int) *ListNode {
    curDummy := &ListNode{}
    cur := curDummy
    tailDummy := &ListNode{}
    tail := tailDummy
    for head != nil {
	if head.Val < x {
	    cur.Next = head
	    cur = cur.Next
	} else {
	    tail.Next = head
	    tail = tail.Next
	}
	head = head.Next
    }
    // 合并两个链表
    tail.Next = nil
    cur.Next = tailDummy.Next
    return curDummy.Next
}

// 重排链表，l1->l2->...->ln-1->ln   =>  l1->ln->l2->ln-1->...
func reorder(order *ListNode) {
    if head == nil {
	return
    }
    // 找中点断开链表
    slow := head
    fast := head.Next
    for fast != nil && fast.Next != nil {
	fast = fast.Next.Next
  	slow = slow.Next
    }
    // 翻转后半部分
    tail := reverse(slow.Next)
    slow.Next = nil
    // 合并链表
    head = mergeReOrderList(head, tail)
}

func reverse(head *ListNode) *ListNode {
    var prev *ListNode
    for head != nil {
	temp := head.Next
	head.Next = prev
	prev = head
	head = temp
    }
    return prev
}

func mergeReOrderList(list1, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
    head := dummy
    toggle := true
    for head != nil {
	if toggle {
	    head.Next = list1
	    list1 = list1.Next
	} else {
	    head.Next = list2
	    list2 = list2.Next
	}
 	toggle = !toggle
	head = head.Next
    }
    if list1 != nil {
 	head.Next = list1
    } else {
	head.Next = list2
    }
    return dummy.Next
}

// 归并排序 排序链表
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
	retyrn head
    }
    // 找中点断链表
    slow := head
    tail := head.Next
    for fast != nil && fast.Next != nil {
	fast = fast.Next.Next
 	slow = slow.Next
    }
    tail := slow.Next
    slow.Next = nil
    // 递归合并
    return mergeSortList(sortList(head), sortList(tail))
}

func mergeSortList(list1, list2 *ListNode) *ListNode {
    dummy := &listNode{}
    head := dummy
    for list1 != nil && list2 != nil {
	if list1.Val < list2.Val {
	    head.Next = list1
	    list1 = list1.Next
	} else {
	    head.Next = list2
	    list2 = list2.Next
	}
	head = head.Next
    }
    if list1 != nil {
	head.Next = list1
    } else {
	head.Next = list2
    }
    return dummy.Next
}

// 环形链表-快慢指针 判断链表是否有环
func hasCycle(head *ListNode) bool {
    if head == nil {
	return false
    }
    fast := head.Next
    slow := head
    for fast != nil && fast.Next != nil {
	if fast == slow {
	    return true
	}
	fast = fast.Next.Next
	slow = slow.Next
    }
    return false
}

// 环形链表-快慢指针 返回链表开始入环的第一个节点
func detectCycle(head *ListNode) *ListNode {
    if head == nil {
	return head
    }
    fast := head.Next
    slow := head
    for fast != nil && fast.Next != nil {
	// 通过指针比较，而不是Val，Val可能存在重复
	if fast == slow {
	    fast = head
	    slow = slow.Next
	    for fast != slow {
		fast = fast.Next
		slow = slow.Next
	    }    
	    return slow
	}
	fast = fast.Next.Next
	slow = slow.Next
    }
    return nil
}

// 判断是否为回文链表
func isPalindrome(head *ListNode) bool {
    if head == nil {
	return true
    }
    // 快慢指针找中点
    slow := head
    fast := head.Next
    for fast != nil && fast.Next != nil {
	fast = fast.Next.Next
	slow = slow.Next 
    }
    // 翻转后半部分
    tail := reverse(slow.Next)
    slow.Next = nil 
    // 前后部分比较判断
    for head != nil && tail != nil {
	if head.Val != tail.Val {
	    return false
	}
  	head = head.Next
	tail = tail.Next
    }
    return true
}

func reverse(head *ListNode) *ListNode {
    var prev *ListNode
    for head != nil {
	temp := head.Next
	head.Next = prev
	prev = head
	head = temp
    }
    return prev
}

// 深拷贝带随机指针的链表
func copyRandomList(head *Node) *Node {
    if head == nil {
   	return head
    }
    for node := head; node != nil; node = node.Next.Next {
	node.Next = &Node{Val: node.Val, Next: node.Next}
    }
    for node := head; node != nil; node = node.Next.Next {
	if node.Random != nil {
	    node.Next.Random = node.Random.Next
	}
    }
    headNew := head.Next
    for node := head; node != nil; node = node.Next {
	nodeNew := node.Next
	node.Next = node.Next.Next
	if nodeNew.Next != nil {
	    nodeNew.Next = nodeNew.Next.Next
	}
    }
    return headNew
}
