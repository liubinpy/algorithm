package basic

import (
    "fmt"
    "testing"
)

func Test_f5(t *testing.T) {
    times := 1000000
    numTimes := make([]int, 37, 37)
    for i := 0; i < times; i++ {
        n := f5()
        numTimes[n-20]++
    }

    for n, t := range numTimes {
        fmt.Printf("得到%d的概率是%f\n", n+20, float64(t)/float64(times))
    }
}
