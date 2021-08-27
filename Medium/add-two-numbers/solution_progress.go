/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 import "math"
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	 
	 l1Current := l1
	 var l1Depth float64 = 0
	 l1Number := 0
	 
	 l2Current := l2
	 var l2Depth float64 = 0
	 l2Number := 0
	 
	 for {
		 if l1Current == nil && l2Current == nil {
			 break
		 }
		 
		 if l1Current != nil {
			 l1Depth++
			 l1Number += l1Current.Val * int(math.Pow(10, l1Depth-1))
			 l1Current = l1Current.Next
		 }
		 
		 if l2Current != nil {
			 l2Depth++
			 l2Number += l2Current.Val * int(math.Pow(10, l2Depth-1))
			 l2Current = l2Current.Next
		 }
	 }
	 
	 addedNumber := l1Number + l2Number
	 
	 var depth float64
	 if l1Depth > l2Depth {
		 depth = l1Depth
	 } else {
		 depth = l2Depth
	 }
	 
	 var result ListNode
	 
	 for i := 0; i < depth; i++ {
		 var nextVal int = addedNumber / int(math.Pow(10, i))
		 addedNumber = addedNumber - (nextVal * int(math.Pow(10, i)))
		 
	 }
	 
	 // for {
	 //     if addedNumber > 0 {
	 //         var nextVal int = addedNumber / int(math.Pow(10, depth-1))
	 //         l3.Val = nextVal
	 //         l4.Next = &l3
	 //         l3 = l4
	 //         addedNumber = addedNumber - (nextVal * int(math.Pow(10, depth-1)))
	 //         depth--
	 //     } else {
	 //         break
	 //     }
	 // }
	 
	 return &l3
 }