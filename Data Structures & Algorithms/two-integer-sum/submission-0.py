class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        elems = {}
        for i, curr in enumerate(nums):
            pair = target - curr
            if pair in elems:
                return [elems[pair], i]
            elems[curr] = i

        