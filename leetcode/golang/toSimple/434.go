package toSimple

import "strings"

func countSegments(s string) int {

	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	split := strings.Split(s, " ")
	if len(split) == 0 {
		return len(s)
	}

	sum := 0

	for _, v := range split {
		if v != "" {
			sum++

		}
	}

	return sum
}
