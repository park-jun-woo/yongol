//ff:func feature=gen-ir type=util control=sequence
//ff:what parseInputValue -- 단일 SSaC Input key-value → FieldArg 파싱

package ir

import "strings"

// parseInputValue converts a single key-value pair from the SSaC Inputs map
// into a FieldArg.
func parseInputValue(key, value string) FieldArg {
	fa := FieldArg{Key: key}

	// Check for quoted literal
	if strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`) && len(value) >= 2 {
		fa.Literal = value[1 : len(value)-1]
		fa.IsQuoted = true
		return fa
	}

	// Check for dotted reference: "source.Field"
	if dotIdx := strings.IndexByte(value, '.'); dotIdx >= 0 {
		fa.Source = value[:dotIdx]
		fa.Field = value[dotIdx+1:]
		return fa
	}

	// Plain variable name or literal
	fa.Source = value
	return fa
}
