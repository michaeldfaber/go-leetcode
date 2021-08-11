func intToRoman(num int) string {
    roman := ""
    
    firstDigit := num % 10
    roman = fmt.Sprintf("%s%s", getRomanCharacters(firstDigit, "I", "V", "X"), roman)
    
    secondDigit := num % 100 - firstDigit
    if num >= 10 {
        roman = fmt.Sprintf("%s%s", getRomanCharacters(secondDigit / 10, "X", "L", "C"), roman)
    }
    
    thirdDigit := num % 1000 - secondDigit - firstDigit
    if num >= 100 {
        roman = fmt.Sprintf("%s%s", getRomanCharacters(thirdDigit / 100, "C", "D", "M"), roman)
    }
    
    fourthDigit := num % 10000 - thirdDigit - secondDigit - firstDigit
    if num >= 1000 {
        roman = fmt.Sprintf("%s%s", getRomanCharacters(fourthDigit / 1000, "M", "", ""), roman)
    }
    
    return roman
}

func getRomanCharacters(digit int, one string, two string, three string) string {
    switch digit {
        case 1:
            return fmt.Sprintf("%s", one)
        case 2:
            return fmt.Sprintf("%s%s", one, one)
        case 3:
            return fmt.Sprintf("%s%s%s", one, one, one)
        case 4:
            return fmt.Sprintf("%s%s", one, two)
        case 5:
            return fmt.Sprintf("%s", two)
        case 6:
            return fmt.Sprintf("%s%s", two, one)
        case 7:
            return fmt.Sprintf("%s%s%s", two, one, one)
        case 8:
            return fmt.Sprintf("%s%s%s%s", two, one, one, one)
        case 9:
            return fmt.Sprintf("%s%s", one, three)
        default:
            return ""
    }
}