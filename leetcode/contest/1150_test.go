package contest

import "testing"

func Test_isMajorityElement(t *testing.T) {
    type args struct {
        nums   []int
        target int
    }
    tests := []struct {
        name string
        args args
        want bool
    }{
        {"1", args{nums: []int{2, 4, 5, 5, 5, 5, 5, 6, 6}, target: 5}, true},
        {"2", args{nums: []int{10, 100, 101, 101}, target: 101}, false},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := isMajorityElement(tt.args.nums, tt.args.target); got != tt.want {
                t.Errorf("isMajorityElement() = %v, want %v", got, tt.want)
            }
        })
    }
}
