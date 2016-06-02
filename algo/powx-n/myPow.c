#define INT_MIN (-1 & 0x80000000)

double myPow(double x, int n) {
    double pow;
    unsigned int np;
    
    if (n < 0) {
        x = 1 / x;
        if (n == INT_MIN) {
            np = (unsigned int) n;
        } else {
            np = -n;
        }
    } else {
        np = n;
    }
    pow = 1.0;
    while (np) {
        if (np & 0x1) {
            pow *= x;
        }
        x *= x;
        np = np >> 1;
    }
    return pow;
}
