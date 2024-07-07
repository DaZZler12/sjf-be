package utils

import "regexp"

// SanitizeString will return a sanitized string, to prevent any kind of injection attacks.
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
