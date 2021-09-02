func maxProfit(prices []int) int {
    
    min := 100001
    profit := 0
    
    for i := 0; i < len(prices); i++ {
        if prices[i] < min {
            min = prices[i]
            continue
        }
        
        if prices[i] - min > profit {
            profit = prices[i] - min
        }
    }
    
    return profit
}