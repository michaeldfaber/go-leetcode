var maxLevel int

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    maxLevel = 0
    getMaxLevel(root, 1)
    return maxLevel
}

func getMaxLevel(root *TreeNode, level int) {
    if root == nil {
        return
    }
    
    if level > maxLevel {
        maxLevel = level
    }
    
    getMaxLevel(root.Left, level + 1)
    getMaxLevel(root.Right, level + 1)
}