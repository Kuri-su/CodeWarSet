package golang

import (
	"reflect"
	"testing"
)

func Test_indexPairs(t *testing.T) {
	type args struct {
		text  string
		words []string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"a", args{text: "ababa", words: []string{"aba", "ab"}}, [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := indexPairs(tt.args.text, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("indexPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
