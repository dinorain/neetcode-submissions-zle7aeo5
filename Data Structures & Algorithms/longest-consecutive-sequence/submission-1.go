func longestConsecutive(nums []int) int {
    numSet := make(map[int]struct{})
    for _, num := range nums {
        numSet[num] = struct{}{}
    }

    longest := 0
    for _, num := range nums {
        if _, ok := numSet[num-1]; !ok {
            length := 1
            for {
                if _, exists := numSet[num+length]; exists {
                    length++
                } else {
                    break
                }
            }
            if length > longest {
                longest = length
            }
        }
    }
    return longest
}