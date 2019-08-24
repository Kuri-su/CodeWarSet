package string

import (
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	fail := "Neither"
	if strings.Contains(IP, ".") {
		// IPV4
		if strings.Contains(IP, "-") {
			return fail
		}
		split := strings.Split(IP, ".")
		if len(split) != 4 {
			return fail
		}
		for _, value := range split {
			if len(value) == 0 || len(value) > 3 || len(value) > 1 && len(value) <= 3 && value[0] == '0' {
				return fail
			}

			i, err := strconv.Atoi(value)
			if err != nil {
				return fail
			} else if i > 255 || i < 0 {
				return fail
			}
		}
		return "IPv4"
	} else {
		// IPV6
		charTable := map[string]bool{
			"1": true,
			"2": true,
			"3": true,
			"4": true,
			"5": true,
			"6": true,
			"7": true,
			"8": true,
			"9": true,
			"0": true,
			"a": true,
			"b": true,
			"c": true,
			"d": true,
			"e": true,
			"f": true,
		}
		split := strings.Split(IP, ":")
		if len(split) != 8 {
			return fail
		}
		for _, value := range split {
			if len(value) > 4 || len(value) == 0 {
				return fail
			}
			for _, char := range value {
				if _, ok := charTable[strings.ToLower(string(char))]; !ok {
					return fail
				}
			}
		}
		return "IPv6"
	}
}
