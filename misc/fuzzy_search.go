package misc

import "strings"

func IsFuzzyMatch(source, target string) bool {
	sourceWords := strings.Split(strings.ToUpper(source), " ")
	targetWords := strings.Split(strings.ToUpper(target), " ")

	if len(sourceWords) == 0 {
		return true
	}
	if len(targetWords) == 0 {
		return false
	}

	for _, v := range sourceWords {
		if !contains(v, targetWords) {
			return false
		}
	}

	return true
}

func contains(e string, a []string) bool {
	for _, v := range a {
		if strings.Contains(v, e) {
			return true
		}
	}
	return false
}
