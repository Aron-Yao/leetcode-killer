func add(result *[]string, str string, m, n int) {
    if m == 0 && n == 0 {
        *result = append(*result, str)
    } else {
        if n > 0 {
            add(result, str + ")", m, n - 1)
        }
        if m > 0 {
            add(result, str + "(", m - 1, n + 1)
        }
    }
}

func generateParenthesis(n int) []string {
    l := &[]string{}
    add(l, "", n, 0)
    return *l
}