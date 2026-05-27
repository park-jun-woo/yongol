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

	// Check for numeric literal before dotted reference (handles "1.5" etc.)
	if isNumericLiteral(value) {
		fa.Literal = value
		return fa
	}

	// Check for dotted reference: "source.Field"
	if dotIdx := strings.IndexByte(value, '.'); dotIdx >= 0 {
		fa.Source = value[:dotIdx]
		fa.Field = value[dotIdx+1:]
		return fa
	}

	// Plain variable name
	fa.Source = value
	return fa
}

// isNumericLiteral returns true when s is a numeric literal: optional leading
// minus, one or more digits, optional single decimal point with digits.
func isNumericLiteral(s string) bool {
	if len(s) == 0 {
		return false
	}
	start := 0
	if s[0] == '-' {
		start = 1
	}
	if start >= len(s) {
		return false
	}
	hasDot := false
	for i := start; i < len(s); i++ {
		if s[i] == '.' && !hasDot {
			hasDot = true
			continue
		}
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}
