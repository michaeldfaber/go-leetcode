func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    
    current := head
    next := head.Next
    
    for {
        if next == nil {
            return head
        }
        
        if current.Val == next.Val {
            if next.Next == nil {
                current.Next = nil
                return head   
            } else {
                current.Next = next.Next
            }
        } else {
            current = next   
        }
        
        next = next.Next
    }
}