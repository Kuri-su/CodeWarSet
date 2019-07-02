package contest

import (
	"reflect"
	"testing"
)

func Test_highFive(t *testing.T) {
	type args struct {
		items [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{[][]int{
			{1, 91}, {1, 92}, {2, 93}, {2, 97}, {1, 60}, {2, 77}, {1, 65}, {1, 87}, {1, 100}, {2, 100}, {2, 76},
		}}, [][]int{
			{1, 87}, {2, 88},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := highFive(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("highFive() = %v, want %v", got, tt.want)
			}
		})
	}
}
