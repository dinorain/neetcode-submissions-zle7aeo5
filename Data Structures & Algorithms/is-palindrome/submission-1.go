func isPalindrome(s string) bool {
    n := len(s)
    if n == 1 {
        return true
    }
    normalized := []rune{}
    for _, c := range s {
        if unicode.IsNumber(c) || unicode.IsLetter(c) {
            normalized = append(normalized, unicode.ToLower(c))
        }
    }

    leftIndex := 0
    rightIndex := len(normalized)-1
    for leftIndex < rightIndex {
        leftChar := normalized[leftIndex]
        rightChar := normalized[rightIndex]
        if leftChar != rightChar {
            return false
        }

        leftIndex++
        rightIndex--
    }

    return true
}