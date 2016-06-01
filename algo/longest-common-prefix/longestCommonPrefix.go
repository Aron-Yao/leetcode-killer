func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    
    end, length := 0, map[int]int{}
    
    for x := 1; x < len(strs); x++ {
        length[x] = len(strs[x])
    }
    
    for y := 0; y < len(strs[0]); y++ {
        saved := strs[0][y]
        for x := 1; x < len(strs); x++ {
            if length[x] <= y || strs[x][y] != saved {
                return strs[0][:end]
            }
        }
        end++
    }
    return strs[0]
}