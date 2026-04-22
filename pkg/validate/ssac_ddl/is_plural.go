//ff:func feature=validate type=util control=sequence topic=ssac-ddl
//ff:what isPlural — 타입명이 이미 복수형인지 판정

package ssac_ddl

import "github.com/jinzhu/inflection"

// isPlural reports whether t is already a plural form. A singular noun is a
// fixed point of Singular() — if Singular(t) differs from t, then t is plural.
// Examples: isPlural("Workflow") = false, isPlural("Workflows") = true.
func isPlural(t string) bool {
	if t == "" {
		return false
	}
	return inflection.Singular(t) != t
}
