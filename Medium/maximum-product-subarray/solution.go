func maxProduct(nums []int) int {
    max := -100001
    product := 0
    
    for i := 0; i < len(nums); i++ {
        product = nums[i]
        if product > max {
            max = product
        }
        for j := i+1; j < len(nums); j++ {            
            product = product * nums[j]
            if product > max {
                max = product
            }
        }
        product = 0
    }
    
    return max
}