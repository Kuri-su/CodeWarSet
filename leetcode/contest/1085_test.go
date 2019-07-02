package contest

import "testing"

func Test_sumOfDigits(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{34, 23, 1, 24, 75, 33, 54, 8}}, 0},
		{"2", args{[]int{99, 77, 33, 66, 55}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfDigits(tt.args.A); got != tt.want {
				t.Errorf("sumOfDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
