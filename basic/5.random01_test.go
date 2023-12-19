package basic

import (
    "fmt"
    "testing"
)

func Test_random1to7(t *testing.T) {
    times := 100000

    numTimes := make([]int, 7, 7)

    for i := 0; i < times; i++ {
        n := random1to7()
        numTimes[n-1] += 1
    }

    for n, t := range numTimes {
        fmt.Printf("得到%d的概率是%f\n", n+1, float64(t)/float64(times))
    }

}
