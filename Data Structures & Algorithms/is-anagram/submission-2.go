func isAnagram(s string, t string) bool {
    nS := len(s)
    nT := len(t)

    if nS != nT {
        return false
    }
    
    charDict := make(map[rune]int, nS)
    for _, charS := range s {
        if _, ok := charDict[charS]; !ok {
            charDict[charS] = 1
            continue
        }
        charDict[charS] += 1
    }

    for _, charT := range t {
        if _, ok := charDict[charT]; !ok {
            return false
        }
        if charDict[charT] == 0 {
            return false
        }
        charDict[charT] -= 1
    }
    return true
}
