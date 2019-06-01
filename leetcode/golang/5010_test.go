package golang

import (
	"testing"
)

func Test_digitsCount(t *testing.T) {
	type args struct {
		d    int
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {"a", args{d: 0, low: 1, high: 200000000}, 1,},
		// {"a", args{d: 1, low: 1, high: 200000000}, 1,},
		// {"a", args{d: 2, low: 1, high: 200000000}, 1,},
		// {"a", args{d: 3, low: 1, high: 200000000}, 1,},
		// {"a", args{d: 4, low: 1, high: 200000000}, 1,},
		// {"a", args{d: 5, low: 1, high: 200000000}, 1,},
		// {"a", args{d: 6, low: 1, high: 200000000}, 1,},
		// {"a", args{d: 7, low: 1, high: 200000000}, 1,},
		// {"a", args{d: 8, low: 1, high: 200000000}, 1,},
		// {"a", args{d: 9, low: 1, high: 200000000}, 1,},
		// {"a", args{d: 1, low: 2, high: 199999999}, 1,},
		// {"a", args{d: 1, low: 2, high: 10}, 1,},
		{"a", args{d: 2, low: 2, high: 199999999}, 1},
		// {"a", args{d: 3, low: 2, high: 199999999}, 1,},
		// {"a", args{d: 4, low: 2, high: 199999999}, 1,},
		// {"a", args{d: 5, low: 2, high: 199999999}, 1,},
		// {"a", args{d: 6, low: 2, high: 199999999}, 1,},
		// {"a", args{d: 7, low: 2, high: 199999999}, 1,},
		// {"a", args{d: 8, low: 2, high: 199999999}, 1,},
		// {"a", args{d: 9, low: 2, high: 199999999}, 1,},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digitsCount(tt.args.d, tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("digitsCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
