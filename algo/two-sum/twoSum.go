import "sort"

type Num struct {
    num int
    index int
}

type Nums []Num

func (n Nums) Less(i, j int) bool {
    return n[i].num < n[j].num
}

func (n Nums) Len() int {
    return len(n)
}

func (n Nums) Swap(i, j int) {
    n[i], n[j] = n[j], n[i]
}


func twoSum(nums []int, target int) []int {
    arr := Nums{}
    var i, j int
    for i, v := range nums {
        arr = append(arr, Num{num: v, index: i})
    }
    
    sort.Sort(arr)
    i, j = 0, len(arr) - 1
    for i < j {
        if arr[i].num + arr[j].num < target {
            i++
        } else if arr[i].num + arr[j].num > target {
            j--
        } else {
            return []int{arr[i].index, arr[j].index}
        }
    }
    return []int{}
}
