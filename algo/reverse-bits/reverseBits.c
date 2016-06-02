uint32_t reverseBits(uint32_t n) {
    uint32_t rvs;
    int cnt;
    
    cnt = 32;
    rvs = 0;
    while (cnt-- > 0) {
        rvs <<= 1;
        rvs |= (n & 0x1);
        n >>= 1;
    }
    return rvs;
}