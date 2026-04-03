func productExceptSelf(nums []int) []int {
    n := len(nums)
    result := make([]int, n)
    result[0] = 1
    for i := 1; i < n; i++ {
        result[i] = nums[i-1] * result[i-1]
    }
    rightProduct := 1
    for i := n - 1; i >= 0; i-- {
        result[i] = result[i] * rightProduct
        rightProduct = rightProduct * nums[i]
    }
    return result
}
