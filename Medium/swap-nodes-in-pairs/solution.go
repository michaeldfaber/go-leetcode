func swapPairs(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    
    if head.Next == nil {
        return head
    }
    
    previous := head
    current := head.Next
    next := current.Next
    
    current.Next = previous
    previous.Next = swapPairs(next)
    
    return current
}