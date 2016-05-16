func isVowels(x byte) bool {
    return x == 'a' || x == 'e' || x == 'i' || x == 'o' || x == 'u' || x == 'A' || x == 'E' || x == 'I' || x == 'O' || x == 'U'
}

func reverseVowels(s string) string {
    slice := []byte(s)
    start, end := 0, len(s) - 1
    for start < end {
        if isVowels(slice[start]) {
            for start < end && !isVowels(slice[end]) {
                end--
            }
            slice[start], slice[end] = slice[end], slice[start]
            end--
        }
        start++
    }
    return string(slice)
}