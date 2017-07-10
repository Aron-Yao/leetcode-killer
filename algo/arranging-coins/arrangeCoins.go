import "math"

func sqrt(n int) int {
    ret := math.Sqrt(float64(n))
    return int(ret)
}

func arrangeCoins(n int) int {
    x := sqrt(2*n)
    for (x*x+x) > 2*n {
        x -= 1
    }a
    
    return x
}
