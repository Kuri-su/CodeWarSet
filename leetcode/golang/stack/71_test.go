package stack

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{path: "/../"}, "/"},
		{"2", args{path: "/home//foo/"}, "/home/foo"},
		{"3", args{path: "/a/./b/../../c/"}, "/c"},
		{"4", args{path: "/a/../../b/../c//.//"}, "/c"},
		{"5", args{path: "/a//b////c/d//././/.."}, "/a/b/c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
