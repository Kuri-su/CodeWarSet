package contest

import (
	"reflect"
	"testing"
)

func Test_permute(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{S: "{a,b}{z,x,y}"}, []string{"ax", "ay", "az", "bx", "by", "bz"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute(tt.args.S); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
