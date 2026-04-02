class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        numsDict = {}
        for i in nums:
            if i in numsDict:
                return True
            numsDict[i] = i
        return False