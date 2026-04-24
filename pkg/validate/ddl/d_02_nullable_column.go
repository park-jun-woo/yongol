//ff:func feature=validate type=rule control=iteration dimension=3 topic=ddl-structural
//ff:what D-2 — NOT NULL missing

package ddl

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// d02NullableColumn validates D-2: a column definition in db/*.sql that lacks
// NOT NULL is an ERROR. PRIMARY KEY columns are implicitly NOT NULL and are
// excluded.
func d02NullableColumn(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	files := readDBSQLFiles(fs.SpecsDir)
	if len(files) == 0 {
		return nil
	}

	// Build a lookup of NullableAnnot maps keyed by table name so we can
	// honor the `-- @nullable` annotation that the parser already captured
	// in fs.DDLTables. Without this lookup the advice message would be
	// misleading: it promises exemption on `-- @nullable` but the rule
	// would still fire. (BUG-028)
	nullableByTable := make(map[string]map[string]bool, len(fs.DDLTables))
	for i := range fs.DDLTables {
		tbl := &fs.DDLTables[i]
		nullableByTable[tbl.Name] = tbl.NullableAnnot
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
				if strings.Contains(upper, "PRIMARY KEY") || strings.Contains(upper, "NOT NULL") {
					continue
				}
				// `-- @nullable` on the column exempts it — the parser
				// already recorded this in Table.NullableAnnot.
				if nullableCols != nil && nullableCols[colName] {
					continue
				}
				diags = append(diags, diagnostic.Diagnostic{
					File:  f.path,
					Line:  blk.startLine + offset,
					Phase: diagnostic.PhaseValidate,
					Level: diagnostic.LevelError,
					Message: fmt.Sprintf(
						"[D-2] table %q column %q — NOT NULL is missing",
						blk.tableName, colName),
					Advice: fmt.Sprintf("Add a NOT NULL constraint to %s.%s, or if intentional add a -- @nullable comment", blk.tableName, colName),
				})
			}
		}
	}
	return diags
}
