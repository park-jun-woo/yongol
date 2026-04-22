//ff:func feature=validate type=rule control=sequence topic=ssac-ddl
//ff:what checkSeqResultType — validates a single sequence's @result type against DDL table existence and singular/plural context

package ssac_ddl

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// checkSeqResultType validates XDS-12 for a single sequence.
// Precondition: seq.Type != "call" && seq.Package == "" && seq.Result != nil.
//
// Two checks:
//   1. Plural element type in singular context — `@get Workflows wf = ...` with
//      Wrapper="" and no []. The plural form is suspect because the wrapper
//      (Page/Cursor/slice) is normally what carries plurality; the element
//      type should be singular. inflection.Plural is idempotent so without
//      this guard the rule below accepts Workflows ≡ workflows in DDL.
//   2. Standard coverage — the sqlc row type or modelToTable(type) must exist as a DDL table.
func checkSeqResultType(fs *yongol.Fullstack, tables map[string]bool, fn ssac.ServiceFunc, seq ssac.Sequence) []diagnostic.Diagnostic {
	typeName := normalizeTypeName(seq.Result.Type)
	if typeName == "" || primitiveTypes[typeName] {
		return nil
	}
	if isSqlcRowType(fs, typeName) {
		return nil
	}
	singularContext := seq.Result.Wrapper == "" && !strings.HasPrefix(seq.Result.Type, "[]")
	if singularContext && isPlural(typeName) {
		return []diagnostic.Diagnostic{{
			File:    fn.FileName,
			Line:    seq.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[XDS-12] @result type %q is plural in a singular context", seq.Result.Type),
			Advice:  "Use the singular form, or wrap in []T / Page[T] / Cursor[T]",
		}}
	}
	tableName := modelToTable(typeName)
	if tables[tableName] {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    fn.FileName,
		Line:    seq.Line,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelWarning,
		Message: fmt.Sprintf("[XDS-12] @result type %q has no matching DDL table %q", seq.Result.Type, tableName),
		Advice:  fmt.Sprintf("Define table %s in the DDL or change the result type", tableName),
	}}
}
