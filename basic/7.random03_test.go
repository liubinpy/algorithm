package basic

import (
    "fmt"
    "testing"
)

func Test_p1(t *testing.T) {
    times := 100000
    count := 0
    for i := 0; i < times; i++ {
        if p2() == 1 {
            count++
        }
    }

    fmt.Printf("得到1得概率是%f\n", float64(count)/float64(times))
}
