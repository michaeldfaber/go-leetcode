func letterCombinations(digits string) []string {
    var letters = [][]string {
        { "a", "b", "c" },
        { "d", "e", "f" },
        { "g", "h", "i" },
        { "j", "k", "l" },
        { "m", "n", "o" },
        { "p", "q", "r", "s" },
        { "t", "u", "v" },
        { "w", "x", "y", "z" },
    }
    
    length := len(digits)
    if length == 0 {
        return []string {}
    }
    if length == 1 {
        digit, _ := strconv.Atoi(digits)
        return letters[digit - 2]
    }
    
    digitsSplice := strings.Split(digits, "")
    firstDigit, _ := strconv.Atoi(digitsSplice[0])
    possibleLetterCombinations := letters[firstDigit - 2]
    
    for i := 1; i < length; i++ {
        newDigit, _ := strconv.Atoi(digitsSplice[i])
        lettersForNewDigit := letters[newDigit - 2]
        var newSplice []string
        
        for _, letterCombination := range possibleLetterCombinations {
            for _, letterForNewDigit := range lettersForNewDigit {
                newSplice = append(newSplice, fmt.Sprintf("%s%s", letterCombination, letterForNewDigit))
            }
        }
        
        possibleLetterCombinations = newSplice
    }
    
    return possibleLetterCombinations
}