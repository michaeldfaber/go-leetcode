func deepestLeavesSum(root *TreeNode) int {
    if root == nil {
        return 0
    }
    _, sum := calculateSum(root, 0, 0, 0)
    return sum
}

func calculateSum(root *TreeNode, level int, maxLevel int, sum int) (int, int) {
    if root == nil {
        return maxLevel, sum
    }
    
    maxLevel, sum = calculateSum(root.Left, level + 1, maxLevel, sum)
    maxLevel, sum = calculateSum(root.Right, level + 1, maxLevel, sum)
    
    if level > maxLevel {
        maxLevel = level
        sum = 0
    }

    if level == maxLevel {
        sum = sum + root.Val
    }
    
    return maxLevel, sum
}