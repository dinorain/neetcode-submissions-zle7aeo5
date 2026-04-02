func twoSum(nums []int, target int) []int {
    numsDict := make(map[int]int, len(nums))
    for i, num := range nums {
        pair := target - num
        if j, ok := numsDict[pair]; ok {
            return []int{j, i}
        }
        numsDict[num] = i
    }
    return []int{}
}
