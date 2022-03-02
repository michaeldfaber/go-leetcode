func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }
    
    if head.Next == nil {
        return false
    }
    
    tortoise := head
    hare := head.Next
    
    for tortoise != hare {
        if hare.Next == nil {
            return false
        }
        
        if hare.Next.Next == nil {
            return false
        }
        
        hare = hare.Next.Next
        tortoise = tortoise.Next
    }
    
    return true
}