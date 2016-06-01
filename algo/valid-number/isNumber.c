typedef enum {
    TOKEN_BLANK = 0,
    TOKEN_SIGN, 
    TOKEN_NUM,
    TOKEN_INT,
    TOKEN_EPX
} TOKEN_KIND;

bool eat(char** s, TOKEN_KIND kind) {
    char* p = *s;
    bool hasNum;
    char* t;
    
    switch (kind) {
        case TOKEN_BLANK:
            while (*p == ' ') {
                ++p;
            }
            if (p == *s) {
                return false;
            }
            break;
        case TOKEN_SIGN:
            if (*p == '+' || *p == '-') {
                ++p;
            } else {
                return false;
            }
            break;            
        case TOKEN_NUM:
            hasNum = false;
            while (*p >= '0' && *p <= '9') {
                ++p;
            }
            if (p != *s) {
                hasNum = true;
            }
            if (*p == '.') {
                ++p;
            }
            t = p;
            while (*p >= '0' && *p <= '9') {
                ++p;
            }
            if (t == p && !hasNum) {
                return false;
            }
            break;
        case TOKEN_INT:
            while (*p >= '0' && *p <= '9') {
                ++p;
            }
            if (p == *s) {
                return false;
            }
            break;
        case TOKEN_EPX:
            if (*p == 'e') {
                ++p;
            } else {
                return false;
            }
            break;
        default:
            return false;
    }
    
    *s = p;
    return true;
}

bool isNumber(char* s) {
    if (s == NULL) {
        return false;
    }
    
    eat(&s, TOKEN_BLANK);
    eat(&s, TOKEN_SIGN);
    if (eat(&s, TOKEN_NUM)) {
        if (eat(&s, TOKEN_EPX)) {
            eat(&s, TOKEN_SIGN);
            if (!eat(&s, TOKEN_INT)) {
                return false;
            }
        }
        eat(&s, TOKEN_BLANK);
        if (*s == '\0') {
            return true;
        }
    }
    return false;
}