package stack

import "testing"

func Test_trap1(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}}, 6},
		{"2", args{height: []int{2, 0, 2}}, 2},
		{"3", args{height: []int{2, 1, 0, 2}}, 3},
		{"4", args{height: []int{2, 0, 1, 0, 2}}, 5},
		{"4", args{height: []int{0, 0, 0, 0, 2}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap1(tt.args.height); got != tt.want {
				t.Errorf("trap1() = %v, want %v", got, tt.want)
			}
		})
	}
}
