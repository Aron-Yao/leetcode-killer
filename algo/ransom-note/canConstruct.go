func canConstruct(ransomNote string, magazine string) bool {
    r := make(map[rune]int)
    m := make(map[rune]int)
    
    for _, c := range ransomNote {
        r[c]++
    }
    
    for _, c := range magazine {
        m[c]++
    }
    
    for k, v := range r {
        t, ok := m[k]
        if !ok {
            return false
        }
        if t < v {
            return false
        } 
    }
    return true
}
