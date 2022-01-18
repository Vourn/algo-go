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
