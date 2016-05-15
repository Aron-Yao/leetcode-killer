#define IS_VOWELS(x) (x == 'a' || x == 'e' || x == 'i' || x == 'o' || x == 'u' || x == 'A' || x == 'E' || x == 'I' || x == 'O' || x == 'U')

#define SWAP(a, b) do { char tmp; tmp = *a; *a = *b; *b = tmp; } while(0)

char* reverseVowels(char* s) {
    char* start;
    char* end;
    
    if (s == NULL) {
        return NULL;
    }
    
    start = s, end = s + strlen(s) - 1;
    while (start < end) {
        if (IS_VOWELS(*start)) {
            while (start < end && !IS_VOWELS(*end)) {
                --end;
            }
            SWAP(start, end);
            --end;
        } 
        ++start;
    }
    return s;
}