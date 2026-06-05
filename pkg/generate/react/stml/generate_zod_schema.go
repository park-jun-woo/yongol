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

	names := make([]string, 0, len(fields))
	for name := range fields {
		names = append(names, name)
	}
	sort.Strings(names)

	var fieldLines []string
	for _, name := range names {
		fc := fields[name]
		fieldLines = append(fieldLines, fmt.Sprintf("  %s: %s,", name, zodChainFor(operationID, name, fc)))
	}

	return fmt.Sprintf("const %s = z.object({\n%s\n})", schemaName, strings.Join(fieldLines, "\n"))
}
