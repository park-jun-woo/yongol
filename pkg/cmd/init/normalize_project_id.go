//ff:func feature=cli-init type=util control=sequence
//ff:what normalizeProjectID — PascalCase/snake_case ProjectID → lowercase snake_case for manifest.metadata.name

package cliinit

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// projectIDPattern restricts ProjectID to alphanumerics and underscores, with
// the first character being a letter. That matches PascalCase ("MyApp") and
// snake_case ("my_app") inputs. Hyphens / dots / path separators are rejected.
var projectIDPattern = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_]*$`)

// ValidateProjectID ensures the CLI positional argument follows the supported
// naming rules. Returns an error message suitable for surfacing to the user.
func ValidateProjectID(id string) error {
	if id == "" {
		return fmt.Errorf("ProjectID is empty")
	}
	if !projectIDPattern.MatchString(id) {
		return fmt.Errorf("ProjectID %q must start with a letter and contain only [A-Za-z0-9_] (PascalCase or snake_case)", id)
	}
	return nil
}

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
		if i > 0 {
			prev := runes[i-1]
			// camelCase boundary: lower|digit → Upper
			if (unicode.IsLower(prev) || unicode.IsDigit(prev)) && unicode.IsUpper(r) {
				out.WriteByte('_')
			}
			// acronym boundary: Upper Upper lower → insert before lower
			if i+1 < len(runes) && unicode.IsUpper(prev) && unicode.IsUpper(r) && unicode.IsLower(runes[i+1]) {
				out.WriteByte('_')
			}
		}
		out.WriteRune(unicode.ToLower(r))
	}
	result := out.String()
	// Collapse repeated underscores and trim edges.
	for strings.Contains(result, "__") {
		result = strings.ReplaceAll(result, "__", "_")
	}
	result = strings.Trim(result, "_")
	return result
}
