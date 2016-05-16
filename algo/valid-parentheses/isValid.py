class Solution(object):
    def isValid(self, s):
        """
        :type s: str
        :rtype: bool
        """
        stack = []
        pmap = {'{' : '}', '(' : ')', '[' : ']'}
        
        for c in s:
            if c in pmap:
                stack.append(pmap[c])
            else:
                if len(stack) == 0:
                    return False
                else:
                    if stack[-1] != c:
                        return False
                    else:
                        stack = stack[:-1]
        return len(stack) == 0