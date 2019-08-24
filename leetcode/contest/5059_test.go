package contest

import "testing"

func Test_calculateTime(t *testing.T) {
	type args struct {
		keyboard string
		word     string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"abcdefghijklmnopqrstuvwxyz", "cba"}, 4},
		{"2", args{"pqrstuvwxyzabcdefghijklmno", "leetcode"}, 73},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateTime(tt.args.keyboard, tt.args.word); got != tt.want {
				t.Errorf("calculateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
