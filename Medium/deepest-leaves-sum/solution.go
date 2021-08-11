var maxLevel int
var sum int

func deepestLeavesSum(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    maxLevel = 0
    sum = 0
    
    calculateSum(root, 0)
    return sum
}

func calculateSum(root *TreeNode, level int) {
    if root == nil {
        return
    }
    
    if level > maxLevel {
        sum = 0
        maxLevel = level
    }

    if level == maxLevel {
        sum = sum + root.Val
    }

    calculateSum(root.Left, level + 1)
    calculateSum(root.Right, level + 1)
}