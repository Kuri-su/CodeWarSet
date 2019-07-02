package contest

import (
	"testing"
)

func Test_numKLenSubstrNoRepeats(t *testing.T) {
	type args struct {
		S string
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"havefunonleetcode", 5}, 6},
		{"2", args{"home", 5}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numKLenSubstrNoRepeats(tt.args.S, tt.args.K); got != tt.want {
				t.Errorf("numKLenSubstrNoRepeats() = %v, want %v", got, tt.want)
			}
		})
	}
}
