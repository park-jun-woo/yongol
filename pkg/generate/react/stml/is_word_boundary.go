//ff:func feature=stml-gen type=util control=sequence
//ff:what PascalCase 룬 배열에서 단어 경계인지 판별한다
package stml

import "unicode"

// isWordBoundary returns true if the rune at index i starts a new word
// in a PascalCase or ALLCAPS sequence.
func isWordBoundary(runes []rune, i int) bool {
	if !unicode.IsUpper(runes[i]) {
		return false
	}
	// Uppercase followed by lowercase → new word (e.g. "Id" in "RoomId")
	if i+1 < len(runes) && unicode.IsLower(runes[i+1]) {
		return true
	}
	// Uppercase preceded by non-uppercase → new word (e.g. start of "ID" after lowercase)
	return !unicode.IsUpper(runes[i-1])
}
