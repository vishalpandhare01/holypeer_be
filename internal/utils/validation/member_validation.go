package validation

import "strings"

func CheckTodysFeelExist(text string) bool {
	text = strings.ToLower(text)
	data := []string{"depression", "anxiety", "relationships", "ocd", "parenting", "family", "loneliness", "happy", "good"}
	for _, val := range data {
		if text == val {
			return true
		}
	}
	return false
}
