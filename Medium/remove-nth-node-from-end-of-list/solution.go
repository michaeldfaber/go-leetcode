func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head.Next == nil {
        return nil
    }
    
    depth := 2
    current := head.Next
    
    for {
        if current.Next != nil {
            current = current.Next
            depth++
        } else {
            break
        }
    }
    
    elementToRemove := depth - n
    
    var previous *ListNode
    current = head
    
    for i := 0; i <= elementToRemove; i++ {
        if i == elementToRemove {
            
            if previous == nil {
                head = head.Next
                return head
            }
            
            previous.Next = current.Next
            continue
        }
        
        previous = current
        current = current.Next
    }
    
    return head
}