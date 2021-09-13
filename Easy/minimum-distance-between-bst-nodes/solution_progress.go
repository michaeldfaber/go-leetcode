func minDiffInBST(root *TreeNode) int {
    min := 100001
    helper(root, &min)
    return min
}

func helper(root *TreeNode, min *int) {
    if root.Left != nil {
        if root.Val - root.Left.Val < *min {
            *min = root.Val - root.Left.Val
        }
        helper(root.Left, min)
    }
    
    if root.Right != nil {
        if root.Right.Val - root.Val < *min {
            *min = root.Right.Val - root.Val
        }
        helper(root.Right, min)
    }
}