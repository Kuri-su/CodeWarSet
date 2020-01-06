package toSimple

import "strings"

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	if len(s) <= 0 {
		return 0
	}
	split := strings.Split(s, " ")
	if len(split) <= 0 {
		return len(s)
	}

	return len(split[len(split)-1])
}
