// O(n)

func maxProfit(prices []int) int {
    
    min := 100001
    max := -10001
    profit := 0
    
    for i := 0; i < len(prices); i++ {
        if prices[i] < min {
            min = prices[i]
            max = -10001
            continue
        }
        
        if prices[i] > max {
            max = prices[i]
            if (max - min) > profit {
                profit = max - min
            }
        }        
    }
    
    return profit
}