int lengthOfLastWord(char* s) {
    char* p;
    int length;
    
    if (s == NULL) {
        return 0;
    }
    
    p = s;
    length = 0;
    while (*p) {
        ++p;
    }
    --p;
    while (p >= s && *p == ' ') {
        --p;
    }
    while (p >= s && *p != ' ') {
        ++length;
        --p;
    }
    return length;
}