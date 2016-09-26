func twoSum(numbers []int, target int) []int {
    var p, q, length int
    
    length = len(numbers)
    
    if length < 2 {
        return []int{}
    }
    
    p, q = 0, length-1
    
    for p < q {
        if numbers[p] + numbers[q] > target {
            q--
        } 
        if numbers[p] + numbers[q] < target {
            p++
        }
        if numbers[p] + numbers[q] == target {
            return []int{p+1, q+1}
        }
    }
    
    return []int{}
}
