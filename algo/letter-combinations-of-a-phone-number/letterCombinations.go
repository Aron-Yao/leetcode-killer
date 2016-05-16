type reduceFunc func([]string, rune) []string

func reduce(fun reduceFunc, str string) []string {
    if len(str) == 0 {
        return []string{}
    }
    
    accu := []string{""}
    for _, x := range str {
        accu = fun(accu, x)
    }
    return accu
}

func fun(accu []string, update rune) []string {
    kvs := map[rune]string{
        '0' : "",
        '1' : "",
        '2' : "abc",
        '3' : "def",
        '4' : "ghi",
        '5' : "jkl",
        '6' : "mno",
        '7' : "pqrs",
        '8' : "tuv",
        '9' : "wxyz",
    }

    ret := []string{}
    for _, str := range accu {
        for _, r := range kvs[update] {
            ret = append(ret, str + string(r))
        }
    }
    return ret
}

func letterCombinations(digits string) []string {
    return reduce(fun, digits)
}