// Kadane's Algorithm

func maxSubArray(nums []int) int {
    currentMax := nums[0]
    max := nums[0]
    
    for i := 1; i < len(nums); i++ {
        if currentMax + nums[i] > nums[i] {
            currentMax = currentMax + nums[i]
        } else {
            currentMax = nums[i]
        }
        
        if currentMax > max {
            max = currentMax
        }
    }
    
    return max
}