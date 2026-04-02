class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        trace = {}
        for i in s:
            if i not in trace:
                trace[i] = 0
            trace[i] += 1
        
        for i in t:
            if i not in trace:
                return False
            elif trace[i] == 0:
                return False
            else:
                trace[i] -= 1
        return True