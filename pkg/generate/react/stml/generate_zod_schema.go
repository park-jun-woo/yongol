//ff:func feature=stml-gen type=generator control=iteration dimension=1
//ff:what OpenAPI requestBody 제약조건에서 zod 스키마 코드 문자열을 생성한다
package stml

import (
	"fmt"
	"sort"
	"strings"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// generateZodSchema produces a zod schema declaration for the given operationId
// using the field constraints from the OpenAPI requestBody. Returns empty string
// when fields is nil or empty.
func generateZodSchema(operationID string, fields map[string]oapiparser.FieldConstraint) string {
	if len(fields) == 0 {
		return ""
	}

	schemaName := toLowerFirst(operationID) + "Schema"

	// Sort field names for deterministic output.
	names := make([]string, 0, len(fields))
	for name := range fields {
		names = append(names, name)
	}
	sort.Strings(names)

	var fieldLines []string
	for _, name := range names {
		fc := fields[name]
		fieldLines = append(fieldLines, fmt.Sprintf("  %s: %s,", name, zodChain(fc)))
	}

	return fmt.Sprintf("const %s = z.object({\n%s\n})", schemaName, strings.Join(fieldLines, "\n"))
}

// zodChain builds a zod validation chain for a single field constraint.
func zodChain(fc oapiparser.FieldConstraint) string {
	// enum takes precedence — z.enum(["a","b"])
	if len(fc.Enum) > 0 {
		quoted := make([]string, len(fc.Enum))
		for i, v := range fc.Enum {
			quoted[i] = fmt.Sprintf(`"%s"`, v)
		}
		base := fmt.Sprintf("z.enum([%s])", strings.Join(quoted, ", "))
		if !fc.Required {
			base += ".optional()"
		}
		return base
	}

	var parts []string

	switch fc.Type {
	case "integer":
		parts = append(parts, "z.number().int()")
	case "number":
		parts = append(parts, "z.number()")
	case "boolean":
		parts = append(parts, "z.boolean()")
	default: // "string" or fallback
		parts = append(parts, "z.string()")
	}

	// string-specific validations
	if fc.Type == "string" || fc.Type == "" {
		if fc.Format == "email" {
			parts = append(parts, ".email()")
		}
		if fc.Required {
			parts = append(parts, ".min(1)")
		}
		if fc.MinLength != nil && *fc.MinLength > 0 {
			// If required already added .min(1) and minLength is 1, skip duplicate.
			if !fc.Required || *fc.MinLength > 1 {
				parts = append(parts, fmt.Sprintf(".min(%d)", *fc.MinLength))
			}
		}
		if fc.MaxLength != nil {
			parts = append(parts, fmt.Sprintf(".max(%d)", *fc.MaxLength))
		}
		if fc.Pattern != "" {
			parts = append(parts, fmt.Sprintf(`.regex(new RegExp("%s"))`, fc.Pattern))
		}
	}

	// numeric validations
	if fc.Type == "integer" || fc.Type == "number" {
		if fc.Minimum != nil {
			parts = append(parts, fmt.Sprintf(".min(%s)", formatFloat(*fc.Minimum)))
		}
		if fc.Maximum != nil {
			parts = append(parts, fmt.Sprintf(".max(%s)", formatFloat(*fc.Maximum)))
		}
	}

	if !fc.Required {
		parts = append(parts, ".optional()")
	}

	return strings.Join(parts, "")
}

// formatFloat formats a float64 for JS output, dropping the decimal point for
// integer values (e.g. 100.0 → "100", 3.14 → "3.14").
func formatFloat(v float64) string {
	if v == float64(int64(v)) {
		return fmt.Sprintf("%d", int64(v))
	}
	return fmt.Sprintf("%g", v)
}
