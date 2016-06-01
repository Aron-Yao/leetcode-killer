func lengthOfLastWord(s string) int {
    x, ret := len(s) - 1, 0
    for x >= 0 && s[x] == ' ' {
        x--
    }
    for x >= 0 && s[x] != ' ' {
        x--
        ret++
    }
    return ret
}