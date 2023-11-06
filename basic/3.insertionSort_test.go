package basic

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		name string
		s    []int64
		want []int64
	}{
		{
			name: "t1",
			s:    []int64{61, 17, 29, 22, 34, 60, 72, 21, 50, 1, 62},
			want: []int64{1, 17, 21, 22, 29, 34, 50, 60, 61, 62, 72},
		},
		{
			name: "t2",
			s:    []int64{5, 4, -1, 10, 1, 2, 6, 9},
			want: []int64{-1, 1, 2, 4, 5, 6, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertionSort(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
