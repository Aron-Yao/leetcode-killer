/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func list2Arr(l *ListNode) []int {
    arr := []int{}
    p := l
    
    for p != nil {
        arr = append(arr, p.Val)
        p = p.Next
    }
    
    return arr
}

func arr2List(arr []int) *ListNode {
    head := &ListNode{}
    tail := head
    
    for _, a := range arr {
        tail.Next = &ListNode{Val: a}
        tail = tail.Next
    }
    
    return head.Next
}

func add(arr1, arr2 []int) []int {
    sum := []int{}
    len1, len2 := len(arr1), len(arr2)
    diff := len1 - len2
    i, j := 0, 0
    carry := true
    prev, next := -1, 0
    
    if diff > 0 {
        i = diff 
        for x := 0; x < i; x++ {
            sum = append(sum, arr1[x])
        }
    } else if diff < 0 {
        j = 0 - diff
        for x := 0; x < j; x++ {
            sum = append(sum, arr2[x])
        }
    } else {
    }
    
    for i < len(arr1) && j < len(arr2) {
        sum = append(sum, arr1[i]+arr2[j])
        i++
        j++
    }
    
    for carry {
        carry = false
        prev, next = -1, 0
        
        for next < len(sum) {
            if sum[next] >= 10 {
                carry = true
                if prev == -1 {
                    tmp := []int{1}
                    tmp = append(tmp, sum...)
                    sum = tmp
                    sum[next+1] -= 10
                } else {
                    sum[prev] += 1
                    sum[next] -= 10
                }
            }
            next += 1
            prev += 1
        }
    }
    
    return sum
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    arr1 := list2Arr(l1)
    arr2 := list2Arr(l2)
    sumArr := add(arr1, arr2)
    return arr2List(sumArr)
}
