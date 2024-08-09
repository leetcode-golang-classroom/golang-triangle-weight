package internal

import (
	"reflect"
	"testing"
)

func TestFindLargestWeight(t *testing.T) {
	type args struct {
		triangle [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sample1",
			args: args{
				triangle: [][]int{
					{1},
					{2, 3},
					{1, 5, 1},
				},
			},
			want: 9,
		},
		{
			name: "sample2",
			args: args{
				triangle: [][]int{
					{5},
					{5, 4},
					{1, 5, 9},
				},
			},
			want: 18,
		},
		{
			name: "sample3",
			args: args{
				triangle: [][]int{
					{5},
					{5, 4},
					{1, 5, 9},
					{4, 9, 1, 3},
				},
			},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindLargestWeight(tt.args.triangle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindLargestWeight(%v) = %v, want %v", tt.args.triangle, got, tt.want)
			}
		})
	}
}
