func countSegments(s string) int {
    var ret int
    var inSeg bool

    for _, n := range s {
        if n != ' ' {
            inSeg = true
        } else {
            if inSeg {
                ret += 1
                inSeg = false
            }
        }
    }

    if inSeg {
        ret += 1
    }

    return ret
}
