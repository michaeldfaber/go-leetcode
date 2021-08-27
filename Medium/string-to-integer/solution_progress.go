func myAtoi(s string) int {
    s = strings.TrimSpace(s)
    strs := strings.Split(s, " ")
    str := strings.Split(strs[0], ".")[0]
    
    if len(str) == 0 {
        return 0
    }
    
    negative := false
    trim := false
    
    if str[0] == '-' {
        negative = true
        trim = true
    }
    
    if str[0] == '+' {
        trim = true
    }
    
    if trim {
        str = str[1:]
    } 
    
    leadingZeros := true
    numbers := "0123456789"
    for i := 0; i < len(str); i++ {
        // fmt.Printf(str)
        // fmt.Printf("\n")
        if strings.Contains(numbers, string(str[i])) {
            if str[i] == '0' && leadingZeros == true {
                // fmt.Printf(string(str[i]))
                str = str[i+1:]
                i--
            } else {
                leadingZeros = false
            }
            continue
        } else {
            str = str[0:i]
            break
        }
    }
    
    fmt.Printf(str)
    
    n, err := strconv.Atoi(str)
    if err != nil {
        return 0
    }
    
    if negative {
        n = n * -1
    }
    
    if n > 2147483647 {
        n = 2147483647
    }
    
    if n < -2147483648 {
        n = -2147483648
    }

    return n
}