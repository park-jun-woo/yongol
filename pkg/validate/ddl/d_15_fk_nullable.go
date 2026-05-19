//ff:func feature=validate type=rule control=iteration dimension=3 topic=ddl-structural
//ff:what D-15 — FK 컬럼에 NOT NULL이 없고 @nullable 어노테이션도 없으면 WARNING

package ddl

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// d15FkNullable validates D-15: a column with a REFERENCES clause that lacks
// NOT NULL is a WARNING. The `-- @nullable` annotation exempts the column,
// signaling that the nullable FK is intentional.
func d15FkNullable(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	files := readDBSQLFiles(fs.SpecsDir)
	if len(files) == 0 {
		return nil
	}

	// Build a lookup of NullableAnnot maps keyed by table name so we can
	// honor the `-- @nullable` annotation.
	nullableByTable := make(map[string]map[string]bool, len(fs.DDLTables))
	for i := range fs.DDLTables {
		tbl := &fs.DDLTables[i]
		annot := make(map[string]bool, len(tbl.Columns))
		for col, c := range tbl.Columns {
			if c.NullableAnnot {
				annot[col] = true
			}
		}
		nullableByTable[tbl.Name] = annot
	}

	var diags []diagnostic.Diagnostic
	for _, f := range files {
		blocks := extractTableBlocks(f)
		for _, blk := range blocks {
			lines := strings.Split(blk.body, "\n")
			nullableCols := nullableByTable[blk.tableName]
			for offset, line := range lines {
				trimmed := strings.TrimSpace(line)
				if isSkippableDDLLine(trimmed) {
					continue
				}
				m := columnNameRe.FindStringSubmatch(trimmed)
				if m == nil {
					continue
				}
				colName := m[1]
				upper := strings.ToUpper(trimmed)
				if !strings.Contains(upper, "REFERENCES") {
					continue
				}
				if strings.Contains(upper, "NOT NULL") {
					continue
				}
				// `-- @nullable` exempts the column.
				if nullableCols != nil && nullableCols[colName] {
					continue
				}
				diags = append(diags, diagnostic.Diagnostic{
					File:  f.path,
					Line:  blk.startLine + offset,
					Phase: diagnostic.PhaseValidate,
					Level: diagnostic.LevelWarning,
					Message: fmt.Sprintf(
						"[D-15] FK 컬럼 %s은 NOT NULL이어야 합니다. 선택적 관계는 NOT NULL DEFAULT 0 sentinel 패턴을 사용하세요. 의도적 nullable이면 -- @nullable을 추가하세요.",
						colName),
					Advice: fmt.Sprintf("Add NOT NULL to %s.%s, use DEFAULT 0 sentinel pattern, or add -- @nullable if intentional", blk.tableName, colName),
				})
			}
		}
	}
	return diags
}
