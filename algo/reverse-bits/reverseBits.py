class Solution(object):
    def reverseBits(self, n):
        """
        :type n: int
        :rtype: int
        """
        rvs, cnt = 0, 32
        while cnt > 0:
            cnt -= 1
            rvs <<= 1
            rvs |= (0x1 & n)
            n >>= 1
        return rvs