func maxSubArray(nums []int) int {
    max := -100001
    sum := 0
    
    for i := 0; i < len(nums); i++ {
        sum = nums[i]
        if sum > max {
            max = sum
        }
        for j := i+1; j < len(nums); j++ {            
            sum = sum + nums[j]
            if sum > max {
                max = sum
            }
        }
        sum = 0
    }
    
    return max
}