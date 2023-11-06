package basic

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {

	tests := []struct {
		name string
		s    []int64
		want []int64
	}{
		{
			name: "t1",
			s:    []int64{5, 4, 10, 1, 2, 6, 9},
			want: []int64{1, 2, 4, 5, 6, 9, 10},
		},
		{
			name: "t2",
			s:    []int64{5, 4, -1, 10, 1, 2, 6, 9},
			want: []int64{-1, 1, 2, 4, 5, 6, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectionSort(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
