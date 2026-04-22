//ff:func feature=validate type=rule control=iteration dimension=3 topic=ddl-structural
//ff:what D-3 — FK DEFAULT 0 sentinel record missing

package ddl

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// d03SentinelMissing validates D-3: when a column is FK + DEFAULT 0, the
// referenced table must have an id=0 sentinel record
// (INSERT INTO ... VALUES (0, ...)). Missing sentinel is an ERROR.
func d03SentinelMissing(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	files := readDBSQLFiles(fs.SpecsDir)
	if len(files) == 0 {
		return nil
	}

	tableContents := allTableContents(files)

	var diags []diagnostic.Diagnostic
	for _, f := range files {
		blocks := extractTableBlocks(f)
		for _, blk := range blocks {
			lines := strings.Split(blk.body, "\n")
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
				if !strings.Contains(upper, "DEFAULT 0") || !strings.Contains(upper, "REFERENCES") {
					continue
				}
				refMatch := referencesRe.FindStringSubmatch(trimmed)
				if refMatch == nil {
					continue
				}
				refTable := refMatch[1]
				refContent, ok := tableContents[refTable]
				if !ok {
					continue
				}
				if hasSentinelInsert(refContent, refTable) {
					continue
				}
				diags = append(diags, diagnostic.Diagnostic{
					File:  f.path,
					Line:  blk.startLine + offset,
					Phase: diagnostic.PhaseValidate,
					Level: diagnostic.LevelError,
					Message: fmt.Sprintf(
						"[D-3] table %q column %q — FK + DEFAULT 0 but referenced table %q has no id=0 sentinel record",
						blk.tableName, colName, refTable),
					Advice: fmt.Sprintf("Add: INSERT INTO %s (id, ...) VALUES (0, ...) ON CONFLICT DO NOTHING;", refTable),
				})
			}
		}
	}
	return diags
}
