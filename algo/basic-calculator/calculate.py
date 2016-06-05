class Solution(object):
    def calculate(self, s):
        """
        :type s: str
        :rtype: int
        """
        nums = []
        tmp = 0
        sign = 1
        i = 0
        while i < len(s):
            c = s[i]
            if c >= '0' and c <= '9':
                num = int(c)
                j = i + 1
                while j < len(s) and s[j] >= '0' and s[j] <= '9':
                    num = num *10 + int(s[j])
                    j += 1
                i = j
                tmp += num * sign
            else:
                if c == '+':
                    sign = 1
                elif c == '-':
                    sign = -1
                elif c == '(':
                    nums.append(tmp)
                    nums.append(sign)
                    tmp, sign = 0, 1
                elif c == ')':
                    tmp = tmp * nums.pop() + nums.pop()
                else:
                    pass
                i += 1
        return tmp
                
                
                
