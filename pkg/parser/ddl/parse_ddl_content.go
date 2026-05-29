//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what parseDDLContent — SQL 텍스트에서 CREATE TABLE 문을 라인 단위로 파싱
package ddl

import "strings"

func parseDDLContent(content string, tables map[string]*Table, file string) {
	lines := strings.Split(content, "\n")
	var currentTable string
	pendingArchived := false
	pendingFuncManaged := false
	for i, line := range lines {
		lineNum := i + 1
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		if isArchivedAnnotation(trimmed) {
			pendingArchived = true
			continue
		}
		if isFuncManagedAnnotation(trimmed) {
			pendingFuncManaged = true
			continue
		}
		currentTable = parseDDLLine(line, currentTable, tables, pendingArchived, pendingFuncManaged, file, lineNum)
		// annotations were consumed by whichever line followed them
		pendingArchived = false
		pendingFuncManaged = false
	}

	// Second pass: collect @sentinel INSERT blocks and attach them to
	// their target tables. Non-annotated top-level INSERTs are silently
	// skipped here; validation (D-9) flags them.
	for _, r := range parseSentinelInserts(content) {
		if !r.Annotated {
			continue
		}
		t := tables[r.Table]
		if t == nil {
			// INSERT targets a table not defined in this directory.
			// Keep permissive at parse time; validation can decide.
			continue
		}
		t.Sentinels = append(t.Sentinels, SentinelInsert{
			SQL:  r.SQL,
			Line: r.StartLine,
			File: file,
		})
	}
}
