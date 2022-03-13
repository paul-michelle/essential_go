package strTools

func ReverseString(str_to_reverse string) string {
	runes_sclice := []rune(str_to_reverse)

	var runes_length int
	runes_length = len(runes_sclice)

	for i, j := 0, runes_length-1; i < runes_length/2; i, j = i+1, j-1 {
		runes_sclice[i], runes_sclice[j] = runes_sclice[j], runes_sclice[i]
	}

	return string(runes_sclice)
}
