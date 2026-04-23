//ff:func feature=migration type=parser control=sequence
//ff:what collectOnAction — "SET NULL" / "NO ACTION" / "CASCADE" 등 액션 표현 수집
package migration

import "strings"

// collectOnAction reads the tokens of an ON DELETE / ON UPDATE action
// (CASCADE / RESTRICT / SET NULL / SET DEFAULT / NO ACTION). Returns
// the uppercase value and the number of tokens consumed.
func collectOnAction(toks []string) (string, int) {
	if len(toks) == 0 {
		return "", 0
	}
	upper := strings.ToUpper(toks[0])
	if upper == "SET" && len(toks) > 1 {
		return "SET " + strings.ToUpper(toks[1]), 2
	}
	if upper == "NO" && len(toks) > 1 && strings.ToUpper(toks[1]) == "ACTION" {
		return "NO ACTION", 2
	}
	return strings.ToUpper(toks[0]), 1
}
