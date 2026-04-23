//ff:func feature=validate type=rule control=sequence topic=ddl-structural
//ff:what d08CheckLine — 한 줄에서 컬럼명 + SERIAL 계열 타입을 추출해 D-8 진단 생성

package ddl

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// d08CheckLine inspects one line of a CREATE TABLE body. If the line
// defines a column whose raw type is SERIAL / BIGSERIAL / SMALLSERIAL,
// it returns a pointer to a D-8 diagnostic, otherwise nil.
func d08CheckLine(f sqlFile, blk tableBlock, line string, offset int) *diagnostic.Diagnostic {
	trimmed := strings.TrimSpace(line)
	if isSkippableDDLLine(trimmed) {
		return nil
	}
	m := columnNameRe.FindStringSubmatch(trimmed)
	if m == nil {
		return nil
	}
	parts := strings.Fields(trimmed)
	if len(parts) < 2 {
		return nil
	}
	rawType := strings.TrimSuffix(parts[1], ",")
	lower := strings.ToLower(rawType)
	if lower != "serial" && lower != "bigserial" && lower != "smallserial" {
		return nil
	}
	colName := m[1]
	return &diagnostic.Diagnostic{
		File:  f.path,
		Line:  blk.startLine + offset,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: fmt.Sprintf(
			"[D-8] Column `%s.%s` uses %s (banned).",
			blk.tableName, colName, strings.ToUpper(rawType)),
		Advice: fmt.Sprintf(
			"Replace with `%s %s PRIMARY KEY`. "+
				"IDENTITY is the SQL standard equivalent (PostgreSQL 10+) and "+
				"avoids standalone sequence management.",
			colName, serialReplacement(lower)),
	}
}
