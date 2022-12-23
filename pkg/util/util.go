package util

import "fmt"

func Pluralize(count int, singular string, plural string) string {
	if count == 1 {
		return singular
	}
	return plural
}

func LogInfo(format string, a ...interface{}) {
	fmt.Printf("\nℹ️ "+format+"\n", a...)
}

func LogSuccess(format string, a ...interface{}) {
	fmt.Printf("\n✅ "+format+"\n", a...)
}

func LogError(a ...interface{}) {
	fmt.Print("❌ ")
	fmt.Println(a...)
}
