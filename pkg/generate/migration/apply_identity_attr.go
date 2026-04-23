//ff:func feature=migration type=parser control=selection
//ff:what applyIdentityAttr — GENERATED [ALWAYS|BY DEFAULT] AS IDENTITY 토큰 해석 후 Column.Identity 설정
package migration

import "strings"

// applyIdentityAttr inspects the token sequence starting at rest[i],
// which must begin with "GENERATED", and attempts to consume a
// GENERATED [ALWAYS|BY DEFAULT] AS IDENTITY clause. It returns the
// number of tokens consumed (0 when rest[i] is GENERATED but the
// trailing tokens do not form a valid IDENTITY phrase). IDENTITY
// columns are implicitly NOT NULL. A conflict with an existing DEFAULT
// is recorded on t.errs.
func applyIdentityAttr(t *Table, col *Column, rest []string, i int) int {
	// Need at least "GENERATED <ALWAYS|BY> AS IDENTITY" (4 tokens)
	if i+3 >= len(rest) {
		return 0
	}
	if strings.ToUpper(rest[i]) != "GENERATED" {
		return 0
	}
	next := strings.ToUpper(rest[i+1])
	var always bool
	var consumed int
	switch {
	case next == "ALWAYS" && i+3 < len(rest) &&
		strings.ToUpper(rest[i+2]) == "AS" &&
		strings.ToUpper(rest[i+3]) == "IDENTITY":
		always = true
		consumed = 4
	case next == "BY" && i+4 < len(rest) &&
		strings.ToUpper(rest[i+2]) == "DEFAULT" &&
		strings.ToUpper(rest[i+3]) == "AS" &&
		strings.ToUpper(rest[i+4]) == "IDENTITY":
		always = false
		consumed = 5
	default:
		return 0
	}

	if col.Default != "" {
		t.errs = append(t.errs,
			"column "+col.Name+": GENERATED AS IDENTITY conflicts with DEFAULT")
	}
	col.Identity = &IdentitySpec{Always: always}
	col.Nullable = false
	return consumed
}
