func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    
    if head.Next == nil {
        return head
    }
    
    var previous *ListNode
    current := head
    next := head.Next
    
    for {
        if current.Val != next.Val {
            if next.Next == nil {
                return head
            }
            
            previous = current
            current = next
            next = next.Next
            continue
        } else {
            duplicate := next.Val
            
            for {
                if next == nil {
                    return head
                }
                
                if next.Next == nil {
                    if previous == nil {
                        return nil
                    } else {
                        previous.Next = nil
                    }
                    return head
                }
                
                next = next.Next
                
                if duplicate != next.Val {
                    
                    if previous == nil {
                        head = next
                    } else {
                        previous.Next = next
                    }
                    
                    current = next
                    next = next.Next
                    
                    if next == nil {
                        return head
                    }
                    
                    break
                }
            }
        }
    }
}