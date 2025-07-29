package common

import "strings"

func CleanInput(text string) []string {
	result := []string{}

	trimmed := strings.Trim(text, " ")

	words := strings.Split(trimmed, " ")

	for _, w := range words {
		cleaned := strings.ReplaceAll(w, " ", "")

		if len(cleaned) != 0 {
			result = append(result, strings.ToLower(cleaned))
		}
	}

	return result
}
