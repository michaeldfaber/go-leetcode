func mySqrt(x int) int {
    if x == 1 {
        return 1
    }
    
    for i := 2; i <= x/2 + 1; i++ {
        if i * i > x {
            return i-1
        }
    }
    
    return 0 // for compiler
}