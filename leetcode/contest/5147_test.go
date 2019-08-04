package contest

import "testing"

func Test_movesToMakeZigzag(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {"1", args{nums: []int{1, 2, 3,},}, 2,},
		// {"2", args{nums: []int{9, 6, 1, 6, 2},}, 4,},
		// {"3", args{nums: []int{1},}, 0,},
		// {"4", args{nums: []int{1, 2},}, 0,},
		// {"5", args{nums: []int{2, 7, 10, 9, 8, 9},}, 4,},
		{"5", args{nums: []int{7, 4, 8, 9, 7, 7, 5}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movesToMakeZigzag(tt.args.nums); got != tt.want {
				t.Errorf("movesToMakeZigzag() = %v, want %v", got, tt.want)
			}
		})
	}
}
