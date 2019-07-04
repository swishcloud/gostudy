package common

func StringLimitLen(str string, maxLen int) string {
	if len(str) > maxLen {
		runes := []rune(str)
		runes = runes[:maxLen]
		return string(runes)

	}
	return str
}
