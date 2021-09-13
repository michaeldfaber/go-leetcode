func plusOne(digits []int) []int {
    n := len(digits) - 1
    
    if digits[n] != 9 {
        digits[n] = digits[n] + 1
        return digits
    }
    
    for {
        if digits[n] == 9 {
            digits[n] = 0
            if n == 0 {
                newDigits := []int { 1 }
                newDigits = append(newDigits, digits...)
                return newDigits
            }
            if digits[n-1] == 9 {
                n = n-1
                continue
            }
            digits[n-1] = digits[n-1] + 1
        } else {
            return digits
        }
    }
}