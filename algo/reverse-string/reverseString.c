#define SWAP(a, b) do { char tmp; tmp = *a; *a = *b; *b = tmp; } while(0)

char* reverseString(char* s) {
    char* start;
    char* end;
    
    if (s == NULL) {
        return NULL;
    }
    
    start = s, end = s + strlen(s) - 1;;
    
    while (start < end) {
        SWAP(start, end);
        ++start;
        --end;
    }
    return s;
}