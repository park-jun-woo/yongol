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
	colType, _ := extractRawType(parts[1:])
	colType = strings.TrimSuffix(colType, ",")

	col := Column{
		Name:          colName,
		RawType:       colType,
		NotNull:       strings.Contains(upper, "NOT NULL") || strings.Contains(upper, "PRIMARY KEY"),
		IsIdentity:    strings.Contains(upper, "AS IDENTITY"),
		NullableAnnot: colNullable,
		Archived:      colArchived,
		Sensitive:     colSensitive,
	}
	if def := extractDefaultString(line); def != "" {
		col.HasDefault = true
		col.DefaultLiteral = def
	}
	applyVarcharLen(&col, colType)
	if strings.Contains(upper, "CHECK") {
		applyCheckEnum(line, &col)
	}
	t.Columns[colName] = col
	t.ColumnOrder = append(t.ColumnOrder, colName)
	applyInlineConstraints(t, upper, colName, parts)
}
