type TOKEN_KIND int

const (
    TOKEN_BLANK TOKEN_KIND = iota
    TOKEN_SIGN
    TOKEN_NUM
    TOKEN_INT
    TOKEN_EPX
)

func eat(s *[]byte, kind TOKEN_KIND) bool {
    if len(*s) == 0 {
        return false
    }
    
    switch kind {
        case TOKEN_BLANK:
            for i, b := range *s {
                if b != ' ' {
                    *s = (*s)[i:]
                    if i == 0 {
                        return false
                    } else {
                        return true
                    }
                }
            }
            *s = []byte{}
            return true
        case TOKEN_SIGN:
            if (*s)[0] == '+' || (*s)[0] == '-' {
                *s = (*s)[1:]
                return true
            } else {
                return false
            }
        case TOKEN_NUM:
            ok1 := eat(s, TOKEN_INT)
            if len(*s) > 0 {
                if (*s)[0] == '.' {
                    *s = (*s)[1:]
                }
            }
            ok2 := eat(s, TOKEN_INT)
            if !ok1 && !ok2 {
                fmt.Println(1)
                return false
            } else {
                return true
            }
        case TOKEN_INT:
            for i, b := range *s {
                if b < '0' || b > '9' {
                    *s = (*s)[i:]
                    if i == 0 {
                        return false
                    } else {
                        return true
                    }
                }
            }
            *s = []byte{}
            return true
        case TOKEN_EPX:
            if (*s)[0] == 'e' {
                *s = (*s)[1:]
                return true
            } else {
                return false
            }
        default:
            return false
    }
}

func isNumber(s string) bool {
    bs := []byte(s)
    eat(&bs, TOKEN_BLANK)
    eat(&bs, TOKEN_SIGN)
    
    if eat(&bs, TOKEN_NUM) {
        if eat(&bs, TOKEN_EPX) {
            eat(&bs, TOKEN_SIGN)
            if !eat(&bs, TOKEN_INT) {
                return false
            }
        }
        eat(&bs, TOKEN_BLANK)
        if len(bs) == 0 {
            return true
        }
    }
    return false
}