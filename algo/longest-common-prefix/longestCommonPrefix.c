#define STRING_NCOPY(des, src, n) do { memcpy(des, src, n); des[n] = '\0'; } while(0)

char* longestCommonPrefix(char** strs, int strsSize) {
    int i, j;
    char* p;
    
    if (strs == NULL || strsSize <= 0) {
        return "";
    }
    for (i = 0; i < strsSize; ++i) {
        if (strs[i] == NULL) {
            return "";
        }
    }
    
    j = 0;
    p = strs[0];
    while (*p) {
        for (i = 1; i < strsSize; ++i) {
            if (strs[i][j] != *p) {
                goto Done;
            }
        }
        ++p;
        ++j;
    }
    
Done:
    p = (char*) malloc(sizeof(char) * (j + 1));
    STRING_NCOPY(p, strs[0], j);
    return p;
}