//ff:func feature=validate type=rule control=sequence topic=ssac-ddl
//ff:what checkSeqResultType — 단일 시퀀스의 @result 타입을 DDL 테이블 존재 여부 + 단수/복수 위치 대조

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
//   2. Standard coverage — sqlc row type 또는 modelToTable(type) 이 DDL 테이블로 존재해야 함.
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
			Advice:  "단수형으로 변경하거나 []T / Page[T] / Cursor[T] 래퍼로 감싸세요",
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
		Advice:  fmt.Sprintf("DDL 에 테이블 %s 를 정의하거나 result 타입을 변경하세요", tableName),
	}}
}
