//ff:func feature=rule type=loader control=sequence
//ff:what populateGoReservedWords — Go 예약어를 Ground.Lookup에 등록
package ground

import "github.com/park-jun-woo/yongol/pkg/rule"

func populateGoReservedWords(g *rule.Ground) {
	g.Lookup["go.reserved"] = rule.StringSet{
		"break": true, "case": true, "chan": true, "const": true, "continue": true,
		"default": true, "defer": true, "else": true, "fallthrough": true, "for": true,
		"func": true, "go": true, "goto": true, "if": true, "import": true,
		"interface": true, "map": true, "package": true, "range": true, "return": true,
		"select": true, "struct": true, "switch": true, "type": true, "var": true,
	}
}
