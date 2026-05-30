//ff:func feature=validate type=util control=sequence topic=openapi-structural
//ff:what o06RequiredLine — O-6 dangling required 의 진단 라인 해석(components 만 LineIndex, inline 은 0)

package openapi

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// o06RequiredLine resolves the diagnostic line for a dangling required name.
// For components schemas (schemaName != "") it prefers the dangling name's
// property line, falling back to the schema declaration line. Inline schemas
// have no LineIndex entry and resolve to line 0.
func o06RequiredLine(fs *yongol.Fullstack, schemaName, name string) int {
	if schemaName == "" || fs.OpenAPILines == nil {
		return 0
	}
	if line := fs.OpenAPILines.SchemaPropertyLine(schemaName, name); line != 0 {
		return line
	}
	return fs.OpenAPILines.SchemaLine(schemaName)
}
