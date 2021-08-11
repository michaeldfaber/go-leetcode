func intToRoman(num int) string {
    roman := ""
    
    firstDigit := num % 10
    switch firstDigit {
        case 1:
            roman = "I"
        case 2:
            roman = "II"
        case 3:
            roman = "III"
        case 4:
            roman = "IV"
        case 5:
            roman = "V"
        case 6:
            roman = "VI"
        case 7:
            roman = "VII"
        case 8:
            roman = "VIII"
        case 9:
            roman = "IX"
    }
    
    secondDigit := num % 100 - firstDigit
    switch secondDigit {
        case 10:
            roman = fmt.Sprintf("%s%s", "X", roman)
        case 20:
            roman = fmt.Sprintf("%s%s", "XX", roman)
        case 30:
            roman = fmt.Sprintf("%s%s", "XXX", roman)
        case 40:
            roman = fmt.Sprintf("%s%s", "XL", roman)
        case 50:
            roman = fmt.Sprintf("%s%s", "L", roman)
        case 60:
            roman = fmt.Sprintf("%s%s", "LX", roman)
        case 70:
            roman = fmt.Sprintf("%s%s", "LXX", roman)
        case 80:
            roman = fmt.Sprintf("%s%s", "LXXX", roman)
        case 90:
            roman = fmt.Sprintf("%s%s", "XC", roman)
    }
    
    thirdDigit := num % 1000 - secondDigit - firstDigit
    switch thirdDigit {
        case 100:
            roman = fmt.Sprintf("%s%s", "C", roman)
        case 200:
            roman = fmt.Sprintf("%s%s", "CC", roman)
        case 300:
            roman = fmt.Sprintf("%s%s", "CCC", roman)
        case 400:
            roman = fmt.Sprintf("%s%s", "CD", roman)
        case 500:
            roman = fmt.Sprintf("%s%s", "D", roman)
        case 600:
            roman = fmt.Sprintf("%s%s", "DC", roman)
        case 700:
            roman = fmt.Sprintf("%s%s", "DCC", roman)
        case 800:
            roman = fmt.Sprintf("%s%s", "DCCC", roman)
        case 900:
            roman = fmt.Sprintf("%s%s", "CM", roman)
    }
    
    fourthDigit := num % 10000 - thirdDigit - secondDigit - firstDigit
    switch fourthDigit {
        case 1000:
            roman = fmt.Sprintf("%s%s", "M", roman)
        case 2000:
            roman = fmt.Sprintf("%s%s", "MM", roman)
        case 3000:
            roman = fmt.Sprintf("%s%s", "MMM", roman)
    }
    
    return roman
}