//ff:func feature=cli-init type=util control=iteration dimension=1
//ff:what NormalizeProjectID — PascalCase/snake_case ProjectID → lowercase snake_case for manifest.metadata.name

package cliinit

import (
	"strings"
	"unicode"
)

// NormalizeProjectID converts a PascalCase or snake_case ProjectID into the
// lowercase snake_case form used as manifest.metadata.name. The rule is:
//  1. Collapse runs of underscores.
//  2. Insert an underscore before each transition from lower→upper (camelCase
//     boundary: "myApp" → "my_app").
//  3. Insert an underscore before a lower/digit that follows a run of
//     uppercase letters (acronym boundary: "APIKey" → "api_key").
//  4. Lowercase the entire string.
//
// The result is always a non-empty lowercase identifier containing only
// [a-z0-9_]; trimming leading/trailing underscores keeps the manifest value
// clean when inputs like "_Foo_" slip through.
func NormalizeProjectID(id string) string {
	var out strings.Builder
	runes := []rune(id)
	for i, r := range runes {
		writeProjectIDBoundary(&out, runes, i, r)
		out.WriteRune(unicode.ToLower(r))
	}
	return collapseUnderscores(out.String())
}
