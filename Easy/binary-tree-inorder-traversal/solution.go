func inorderTraversal(root *TreeNode) []int {
    result := []int {}
    helper(root, &result)
    return result
}

func helper(root *TreeNode, result *[]int) {
    if root != nil {
        helper(root.Left, result)
        *result = append(*result, root.Val)
        helper(root.Right, result)
    }
}