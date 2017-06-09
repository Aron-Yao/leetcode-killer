func reverse(bs []byte, start, end int) {
    if end >= len(bs) {
        end = len(bs) - 1
    }
    
    for start < end {
        bs[start], bs[end] = bs[end], bs[start]
        start++
        end--
    }
}

func reverseStr(s string, k int) string {
    start := 0
    retStr := []byte(s)
    
    for start < len(s) { 
        reverse(retStr, start, start+k-1)
        start += 2 * k
    }
    
    return string(retStr)
}
