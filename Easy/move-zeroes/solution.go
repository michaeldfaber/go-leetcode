func moveZeroes(nums []int)  {
    i := 0
    z := 0
    
    for x := 0; x < len(nums); x++ {
        if nums[x] != 0 {
            nums[i] = nums[x]
            i++
        } else {
            z++
        }
    }
    
    for y := len(nums)-z; y < len(nums); y++ {
        nums[y] = 0
    }
}