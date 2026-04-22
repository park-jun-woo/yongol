//ff:func feature=manifest type=parser control=sequence
//ff:what parseColumnDef — 컬럼 정의 라인에서 이름, 타입, 인라인 제약 추출
package ddl

import "strings"

func parseColumnDef(line, upper string, t *Table, pendingArchived bool) {
	// strip trailing `-- ...` SQL comment before tokenizing; detect `@archived`.
	colArchived := pendingArchived
	colSensitive := false
	colNullable := false
	if idx := strings.Index(line, "--"); idx >= 0 {
		comment := strings.TrimSpace(line[idx+2:])
		if comment == "@archived" {
			colArchived = true
		}
		if comment == "@sensitive" {
			colSensitive = true
		}
		if comment == "@nullable" {
			colNullable = true
		}
		line = strings.TrimSpace(line[:idx])
		if line == "" {
			return
		}
		upper = strings.ToUpper(line)
	}

	parts := strings.Fields(line)
	if len(parts) < 2 {
		return
	}
	colName := parts[0]
	colType := strings.ToUpper(parts[1])
	colType = strings.TrimSuffix(colType, ",")

	t.Columns[colName] = pgTypeToGo(colType)
	t.ColumnOrder = append(t.ColumnOrder, colName)
	applyInlineConstraints(t, upper, colName, parts)
	applyVarcharLen(t, colName, colType)
	if strings.Contains(upper, "CHECK") {
		applyCheckEnum(line, colName, t)
	}
	if colArchived {
		if t.ArchivedColumns == nil {
			t.ArchivedColumns = make(map[string]bool)
		}
		t.ArchivedColumns[colName] = true
	}
	if colSensitive {
		if t.SensitiveColumns == nil {
			t.SensitiveColumns = make(map[string]bool)
		}
		t.SensitiveColumns[colName] = true
	}
	if def := extractDefaultString(line); def != "" {
		if t.Defaults == nil {
			t.Defaults = make(map[string]string)
		}
		t.Defaults[colName] = def
	}
	if strings.Contains(upper, "NOT NULL") || strings.Contains(upper, "PRIMARY KEY") {
		if t.NotNullCols == nil {
			t.NotNullCols = make(map[string]bool)
		}
		t.NotNullCols[colName] = true
	}
	if colNullable {
		if t.NullableAnnot == nil {
			t.NullableAnnot = make(map[string]bool)
		}
		t.NullableAnnot[colName] = true
	}
}
