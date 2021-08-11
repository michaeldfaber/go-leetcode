func removeElement(nums []int, val int) int {
    i := 0
    for x := 0; x < len(nums); x++ {
        if nums[x] != val {
            nums[i] = nums[x]
            i++
        }
	}
    return i
}