func hasDuplicate(nums []int) bool {
    n := len(nums)
    numDict := make(map[int]bool, n)
    for _, num := range nums {
        if numDict[num] {
            return true
        }
        numDict[num] = true
    }
    return false
}
