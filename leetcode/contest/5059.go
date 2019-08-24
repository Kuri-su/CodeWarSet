package contest

func calculateTime(keyboard string, word string) int {
	charTable := make(map[int32]int)
	for key, value := range keyboard {
		charTable[value] = key
	}

	lastPosition := 0
	sum := 0
	for _, value := range word {
		a := lastPosition - charTable[value]
		lastPosition = charTable[value]
		if a < 0 {
			a *= -1
		}
		sum += a
	}
	return sum
}
