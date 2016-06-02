import "sort"

type V2I struct {
    Value int
    Index int
}

type V2IS []V2I

func (vi V2IS) Len() int {
    return len(vi)
}

func (vi V2IS) Less(a, b int) bool {
    return vi[a].Value < vi[b].Value
}

func (vi V2IS) Swap(a, b int) {
    vi[a], vi[b] = vi[b], vi[a]
}

func abs(a, b int) int {
    if a > b {
        return a - b
    } else {
        return b - a
    }
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
    v2is := []V2I{}
    for i, e := range nums {
        v := V2I{e, i}
        v2is = append(v2is, v)
    }
    sort.Sort(V2IS(v2is))
    for i, e := range v2is {
        for j := i + 1; j < len(v2is); j++ {
            diff := abs(e.Value, v2is[j].Value)
            if diff <= t {
                if abs(e.Index, v2is[j].Index) <= k {
                    return true
                }
            } else {
                break
            }
        }
    }
    return false
}
