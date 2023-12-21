package basic

import (
	"fmt"
	"testing"
)

func Test_binarySearch(t *testing.T) {
	fmt.Println(binarySearch([]int64{1, 2, 2, 3, 6, 7, 11, 12, 13, 19, 22, 25, 26}, 14))
}
