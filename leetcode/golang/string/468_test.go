package string

import "testing"

func Test_validIPAddress(t *testing.T) {
	type args struct {
		IP string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{IP: "172.16.254.1"}, "IPv4"},
		{"2", args{IP: "172.16.254.01"}, "Neither"},
		{"3", args{IP: "2001:0db8:85a3:0000:0000:8a2e:0370:7334"}, "IPv6"},
		{"4", args{IP: "2001:db8:85a3:0:0:8A2E:0370:7334"}, "IPv6"},
		{"5", args{IP: "2001:0db8:85a3::8A2E:0370:7334"}, "Neither"},
		{"6", args{IP: "02001:0db8:85a3:0000:0000:8a2e:0370:7334"}, "Neither"},
		{"7", args{IP: "256.256.256.256"}, "Neither"},
		{"8", args{IP: "20EE:FGb8:85a3:0:0:8A2E:0370:7334"}, "Neither"},
		{"9", args{IP: "20EE:Fb8:85a3:0:0:8A2E:0370:7334"}, "IPv6"},
		{"10", args{IP: "2001:db8:85a3:0::8a2E:0370:7334"}, "Neither"},
		{"11", args{IP: "15.16.-0.1"}, "Neither"},
		{"12", args{IP: "192.0.0.1"}, "IPv4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validIPAddress(tt.args.IP); got != tt.want {
				t.Errorf("validIPAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
