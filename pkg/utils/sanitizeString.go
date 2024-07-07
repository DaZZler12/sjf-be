package utils

import "regexp"

// SanitizeString
//   - removes special characters from the input string
//   - returns the sanitized string
//   - input: string
//   - output: string
//   - example: SanitizeString("Hello, World!") => "Hello World"
func SanitizeString(input string) string {
	if input == "" {
		return ""
	}
	str := regexp.MustCompile(`([.*+?^=!:${}()|\[\]\/\\])`)
	return str.ReplaceAllString(input, "")
}
