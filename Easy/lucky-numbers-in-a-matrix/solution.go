func luckyNumbers (matrix [][]int) []int {
    rowLength := len(matrix[0])
    columnLength := len(matrix)
    
    min := 100001
    max := 0
    
    var minNumbers []int
    var luckyNumbers []int
    
    for i := 0; i < columnLength; i++ {
        for j := 0; j < rowLength; j++ {
            
            if matrix[i][j] < min {
                min = matrix[i][j]
            }
        }
        minNumbers = append(minNumbers, min)
        min = 100001
    }
    
    for i := 0; i < rowLength; i++ {
        for j := 0; j < columnLength; j++ {
            if matrix[j][i] > max {
                max = matrix[j][i]
            }
        }
        
        for _, minNumber := range minNumbers {
            if minNumber == max {
                luckyNumbers = append(luckyNumbers, max)
            }
        }
        max = 0
    }
    
    return luckyNumbers
}