package contest

import "testing"

func Test_assignBikes(t *testing.T) {
	type args struct {
		workers [][]int
		bikes   [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"a", args{workers: [][]int{{0, 0}, {2, 1}}, bikes: [][]int{{1, 2}, {3, 3}}}, 6},
		{"a", args{workers: [][]int{{0, 0}, {1, 1}, {2, 0}}, bikes: [][]int{{1, 0}, {2, 2}, {2, 1}}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := assignBikes(tt.args.workers, tt.args.bikes); got != tt.want {
				t.Errorf("assignBikes() = %v, want %v", got, tt.want)
			}
		})
	}
}
