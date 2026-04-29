//ff:func feature=validate type=util control=sequence topic=ssac-sqlc
//ff:what indexCaseFold — case-insensitive substring 위치 (RETURNING `AS` 분할용)

package ssac_sqlc

import "strings"

// indexCaseFold reports the case-insensitive index of needle in haystack.
// Returns -1 when not found. Used by splitReturningColumns to locate the
// ` AS ` keyword regardless of casing.
func indexCaseFold(haystack, needle string) int {
	return strings.Index(strings.ToUpper(haystack), strings.ToUpper(needle))
}
