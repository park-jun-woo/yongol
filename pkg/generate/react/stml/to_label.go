//ff:func feature=stml-gen type=util control=sequence
//ff:what 필드명(snake_case/PascalCase)을 Title Case 라벨로 변환한다
package stml

import "strings"

// toLabel converts a field name to a human-readable label.
// snake_case → Title Case (e.g., "trigger_event" → "Trigger Event").
// PascalCase → spaced words (e.g., "RoomID" → "Room ID").
func toLabel(name string) string {
	if strings.Contains(name, "_") {
		return snakeToLabel(name)
	}
	return pascalToLabel(name)
}
