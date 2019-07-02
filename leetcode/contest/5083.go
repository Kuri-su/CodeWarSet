package contest

import "strings"

func findOcurrences(text string, first string, second string) []string {
	arr := make([]string, 0)
	fields := strings.Fields(text)
	for key, value := range fields {
		if key >= len(fields)-2 {
			break
		}
		if value == first {
			if fields[key+1] == second {
				arr = append(arr, fields[key+2])
			}
		}
	}
	return arr
}
