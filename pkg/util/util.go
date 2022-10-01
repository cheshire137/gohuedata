package util

func Pluralize(count int, singular string, plural string) string {
	if count == 1 {
		return singular
	}
	return plural
}
