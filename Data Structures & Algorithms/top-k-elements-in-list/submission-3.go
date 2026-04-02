func topKFrequent(nums []int, k int) []int {
    counterDict := make(map[int]int)
    for _, num := range nums {
        counterDict[num]++
    }

    buckets := make([][]int, len(nums)+1)
    for num, count := range counterDict {
        buckets[count] = append(buckets[count], num)
    }

    result := make([]int, 0, k)
    for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
        result = append(result, buckets[i]...)
    }
    return result
}