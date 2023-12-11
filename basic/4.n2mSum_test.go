package basic

import "testing"

func Test_n2mSumWay1(t *testing.T) {
	type args struct {
		arr []int64
		n   int64
		m   int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "t1",
			args: struct {
				arr []int64
				n   int64
				m   int64
			}{arr: []int64{5, 4, -1, 10, 1, 2, 6, 9}, n: 2, m: 7},
			want: 27,
		},
		{
			name: "t2",
			args: struct {
				arr []int64
				n   int64
				m   int64
			}{arr: []int64{5, 4, -1, 10, 1, 2, 6, 9}, n: 6, m: 6},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := n2mSumWay1(tt.args.arr, tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("n2mSumWay1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_n2mSumWay2(t *testing.T) {
	type args struct {
		arr []int64
		n   int64
		m   int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "t1",
			args: struct {
				arr []int64
				n   int64
				m   int64
			}{arr: []int64{5, 4, -1, 10, 1, 2, 6, 9}, n: 2, m: 7},
			want: 27,
		},
		{
			name: "t2",
			args: struct {
				arr []int64
				n   int64
				m   int64
			}{arr: []int64{5, 4, -1, 10, 1, 2, 6, 9}, n: 6, m: 6},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := n2mSumWay2(tt.args.arr, tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("n2mSumWay2() = %v, want %v", got, tt.want)
			}
		})
	}

}
