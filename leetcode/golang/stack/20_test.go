package stack

import (
	"testing"
)

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{s: "()"}, true},
		{"2", args{s: "()[]{}"}, true},
		{"3", args{s: "(]"}, false},
		{"4", args{s: "([)]"}, false},
		{"5", args{s: "{[]}"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid_20(tt.args.s); got != tt.want {
				t.Errorf("isValid_20() = %v, want %v", got, tt.want)
			}
		})
	}
}
