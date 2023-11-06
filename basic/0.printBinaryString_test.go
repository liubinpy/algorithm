package basic

import (
	"math"
	"testing"
)

func TestPrintBinaryString(t *testing.T) {
	testCases := []struct {
		name string
		num  int32
		want string
	}{
		{
			name: "t1",
			num:  20,
			want: "00000000000000000000000000010100",
		},
		{
			name: "t2",
			num:  -1,
			want: "11111111111111111111111111111111",
		},
		{
			name: "t3",
			num:  math.MaxInt32,
			want: "01111111111111111111111111111111",
		},
		{
			name: "t4",
			num:  math.MinInt32,
			want: "10000000000000000000000000000000",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintBinaryString(tt.num); got != tt.want {
				t.Errorf("PrintBinaryString() = %v, want %v", got, tt.want)
			}
		})
	}
}
