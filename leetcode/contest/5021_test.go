package contest

import "testing"

func Test_twoSumLessThanK(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{34, 23, 1, 24, 75, 33, 54, 8}, 60}, 58},
		{"2", args{[]int{10, 20, 30}, 15}, -1},
		{"3", args{[]int{10, 20, 30}, 15}, 198},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumLessThanK(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("twoSumLessThanK() = %v, want %v", got, tt.want)
			}
		})
	}
}
