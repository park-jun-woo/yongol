//ff:func feature=migration type=parser control=iteration dimension=1
//ff:what applyRefOnActions — ON DELETE / ON UPDATE 액션 구문 순회 수집
package migration

import "strings"

// applyRefOnActions walks the "ON DELETE X" / "ON UPDATE Y" tail and
// writes the action values onto fk. Returns the new reader index.
func applyRefOnActions(fk *ForeignKey, toks []string, consumed int) int {
	for consumed+2 < len(toks) {
		if strings.ToUpper(toks[consumed]) != "ON" {
			break
		}
		action := strings.ToUpper(toks[consumed+1])
		val, step := collectOnAction(toks[consumed+2:])
		setFKAction(fk, action, val)
		consumed += 2 + step
	}
	return consumed
}
