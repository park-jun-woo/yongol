//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what scanImports — 비민감 컬럼 타입에 따라 time / json.RawMessage import 필요 여부 판정

package sqlcpost

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// scanImports returns (needsTime, needsJSON) based on the non-sensitive
// column types we'll reference. Sensitive columns are always emitted as
// string literals so they never trigger imports.
func scanImports(t ddl.Table, cols []string) (bool, bool) {
	needsTime, needsJSON := false, false
	for _, col := range cols {
		if t.SensitiveColumns[col] {
			continue
		}
		switch t.Columns[col] {
		case "time.Time":
			needsTime = true
		case "json.RawMessage":
			needsJSON = true
		}
	}
	return needsTime, needsJSON
}
