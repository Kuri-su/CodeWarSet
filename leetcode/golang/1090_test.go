package golang

import "testing"

func Test_largestValsFromLabels(t *testing.T) {
	type args struct {
		values     []int
		labels     []int
		num_wanted int
		use_limit  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {"1", args{[]int{5, 4, 3, 2, 1}, []int{1, 1, 2, 2, 3}, 3, 1}, 9,},
		// {"2", args{[]int{5, 4, 3, 2, 1}, []int{1, 3, 3, 3, 2}, 3, 2}, 12,},
		// {"3", args{[]int{9, 8, 8, 7, 6}, []int{0, 0, 0, 1, 1}, 3, 1}, 16,},
		// {"4", args{[]int{9, 8, 8, 7, 6}, []int{0, 0, 0, 1, 1}, 3, 2}, 24,},
		{"4", args{[]int{2, 6, 1, 2, 6}, []int{2, 2, 2, 1, 1}, 1, 1}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestValsFromLabels(tt.args.values, tt.args.labels, tt.args.num_wanted, tt.args.use_limit); got != tt.want {
				t.Errorf("largestValsFromLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
