func fizzBuzz(n int) []string {
    var results []string
    
    for i := 1; i <= n; i++ {
        result := ""
        
        if i % 3 == 0 {
            result = fmt.Sprintf("%s%s", result, "Fizz")
        }
        
        if i % 5 == 0 {
            result = fmt.Sprintf("%s%s", result, "Buzz")
        }
        
        if result == "" {
            results = append(results, strconv.Itoa(i))
        } else {
            results = append(results, result)
        }
    }
    
    return results
}