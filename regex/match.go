package regex

func MatchChinese(str string) []string {
	return match(patternChinese, str)
}