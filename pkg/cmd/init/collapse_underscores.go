//ff:func feature=cli-init type=util control=iteration dimension=1
//ff:what collapseUnderscores — collapse runs of "__" into "_" and trim edges

package cliinit

import "strings"

// collapseUnderscores collapses repeated underscores in s and trims any
// leading/trailing separators. Used by NormalizeProjectID to keep manifest
// identifiers clean when inputs like "_Foo__Bar_" slip through.
func collapseUnderscores(s string) string {
	for strings.Contains(s, "__") {
		s = strings.ReplaceAll(s, "__", "_")
	}
	return strings.Trim(s, "_")
}
