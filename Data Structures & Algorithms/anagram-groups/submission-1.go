func groupAnagrams(strs []string) [][]string {
    n := len(strs)
    anaDict := make(map[string][]string, n)
    for _, str := range strs {
        strDict := make(map[rune]int, len(str))
        for _, ch := range str {
            strDict[ch]++
        }
        key := fmt.Sprintf("%v", strDict)
        anaDict[key] = append(anaDict[key], str)
    }

    result := make([][]string, 0, len(anaDict))
    for _, group := range anaDict {
        result = append(result, group)
    }
    return result
}
