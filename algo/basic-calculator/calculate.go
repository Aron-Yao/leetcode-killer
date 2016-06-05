func calculate(s string) int {
    nums := []int{}
    tmp, sign := 0, 1
    i := 0
    for i < len(s) {
        c := s[i]
        if c >= '0' && c <= '9' {
            num := int(c - '0')
            j := i + 1
            for j < len(s) && s[j] >= '0' && s[j] <= '9' {
                num = num * 10 + int(s[j] - '0')
                j++
            }
            i = j
            tmp += sign * num
        } else {
            if c == '+' {
                sign = 1
            } else if c == '-' {
                sign = -1
            } else if c == '(' {
                nums = append(nums, tmp, sign)
                tmp, sign = 0, 1
            } else if c == ')' {
                tmp = tmp * nums[len(nums) - 1] + nums[len(nums) - 2]
                nums = nums[:len(nums) - 2]
            } else {
            }
            i++                
        }
    }
    return tmp
}
