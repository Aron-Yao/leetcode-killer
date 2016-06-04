int calculate(char* s) {
    int nums[2];
    char oprs[2];
    int n, m, num;
    char* p;
    char* q;
    
    n = -1, m = -1;
    p = s;
    while (*p) {
        if (*p == ' ') {
            ++p;
        }
        if (*p >= '0' && *p <= '9') {
            num = *p - '0';
            q = p + 1;
            while (*q && (*q >= '0' && *q <= '9')) {
                num *= 10;
                num += *q - '0';
                ++q;
            }
            p = q;
            if (m >= 0 && (oprs[m] == '*' || oprs[m] == '/')) {
                nums[n] = (oprs[m] == '*') ? (nums[n] * num) : (nums[n] / num);
                --m;
            } else {
                nums[++n] = num;
            }
        } else if (*p == '+' || *p == '-' || *p == '*' || *p == '/') {
            if (m >= 0 && (*p == '+' || *p == '-') && (oprs[m] == '+' || oprs[m] == '-')) {
                nums[0] = (oprs[m] == '+') ? (nums[0] + nums[1]) : (nums[0] - nums[1]);
                n = 0;
                --m;
            }
            ++m;
            oprs[m] = *p;
            ++p;
        }
    }
    
    if (m < 0) {
        return nums[0];
    } else {
        return (oprs[m] == '+') ? (nums[0] + nums[1]) : (nums[0] - nums[1]);
    }
}
