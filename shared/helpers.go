package shared

import "strings"

func FindSubstring(slice []string, substring string) string {
	for _, s := range slice {
		if strings.Contains(s, substring) {
			return s
		}
	}
	return "" // Return an empty string if no match is found
}
