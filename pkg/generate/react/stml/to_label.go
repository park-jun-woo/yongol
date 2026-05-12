//ff:func feature=stml-gen type=util control=sequence
//ff:what 필드명(snake_case/PascalCase)을 Title Case 라벨로 변환한다
package stml

import (
	"strings"
	"unicode"
)

// toLabel converts a field name to a human-readable label.
// snake_case → Title Case (e.g., "trigger_event" → "Trigger Event").
// PascalCase → spaced words (e.g., "RoomID" → "Room ID").
func toLabel(name string) string {
	if strings.Contains(name, "_") {
		return snakeToLabel(name)
	}
	return pascalToLabel(name)
}

func snakeToLabel(name string) string {
	parts := strings.Split(name, "_")
	for i, p := range parts {
		if len(p) > 0 {
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
		}
	}
	return strings.Join(parts, " ")
}

func pascalToLabel(name string) string {
	if name == "" {
		return ""
	}
	runes := []rune(name)
	// Capitalize first letter
	runes[0] = unicode.ToUpper(runes[0])
	var words []string
	start := 0
	for i := 1; i < len(runes); i++ {
		if unicode.IsUpper(runes[i]) {
			// Check if this starts a new word or continues an acronym
			if i+1 < len(runes) && unicode.IsLower(runes[i+1]) {
				words = append(words, string(runes[start:i]))
				start = i
			} else if !unicode.IsUpper(runes[i-1]) {
				words = append(words, string(runes[start:i]))
				start = i
			}
		}
	}
	words = append(words, string(runes[start:]))
	return strings.Join(words, " ")
}
