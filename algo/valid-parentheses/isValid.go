func isValid(s string) bool {
    pmap := map[rune]rune{'[' : ']', '(' : ')', '{' : '}'}
    stack := []rune{}

    for _, r := range s {
        if v, ok := pmap[r]; ok {
            stack = append(stack, v)
        } else {
            if len(stack) == 0 {
                return false
            } else {
                last := len(stack) - 1
                if stack[last] == r {
                    stack = stack[:last]
                } else {
                    return false
                }
            }
        }
    }    
    return len(stack) == 0
}