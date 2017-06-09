func reverseWords(s string) string {
    bs := []byte(s)
    now, start, end := 0, 0, 0
    
    for now < len(s) {
        start, end = now, now
        for end < len(s) && bs[end] != ' ' {
            end++
        }
        
        now = end + 1
        end--
        for start < end {
            bs[start], bs[end] = bs[end], bs[start]
            start++
            end--
        }
    }    
    
    return string(bs)
}
