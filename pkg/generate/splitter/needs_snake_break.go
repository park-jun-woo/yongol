//ff:func feature=gen-splitter type=util control=sequence
//ff:what needsSnakeBreak — snake() 에서 i 위치 직전에 "_" 를 삽입할지 결정
package splitter

import "unicode"

// needsSnakeBreak returns true when position i of a rune slice starts a
// new word boundary under the CamelCase → snake_case convention used
// throughout yongol. The rules are:
//
//   - the rune at i must be uppercase;
//   - either the preceding rune is lowercase or a digit (lower→upper
//     boundary), or we are in the middle of an acronym that ends at i
//     (upper→upper→lower sequence — split before the terminal upper).
func needsSnakeBreak(runes []rune, i int) bool {
	if i == 0 || !unicode.IsUpper(runes[i]) {
		return false
	}
	prev := runes[i-1]
	if unicode.IsLower(prev) || unicode.IsDigit(prev) {
		return true
	}
	if !unicode.IsUpper(prev) {
		return false
	}
	if i+1 >= len(runes) {
		return false
	}
	return unicode.IsLower(runes[i+1])
}
