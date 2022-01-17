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
