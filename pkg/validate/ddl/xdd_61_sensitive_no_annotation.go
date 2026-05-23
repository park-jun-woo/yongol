//ff:func feature=validate type=rule control=iteration dimension=3 topic=ddl-structural
//ff:what XDD-61 — sensitive-pattern column missing @sensitive annotation

package ddl

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xdd61SensitiveNoAnnotation validates XDD-61: a column name in db/*.sql that
// matches a sensitive pattern (password, token, ssn, etc.) but lacks a
// `-- @sensitive` (or `-- @nosensitive`) annotation on the same line is a WARNING.
// Without the annotation the column value may be exposed in JSON responses.
func xdd61SensitiveNoAnnotation(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	files := readDBSQLFiles(fs.SpecsDir)
	if len(files) == 0 {
		return nil
	}

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
				match := matchSensitivePattern(colName)
				if match == "" {
					continue
				}
				if hasInlineSensitiveAnnotation(line) {
					continue
				}
				diags = append(diags, diagnostic.Diagnostic{
					File:  f.path,
					Line:  blk.startLine + offset,
					Phase: diagnostic.PhaseValidate,
					Level: diagnostic.LevelWarning,
					Message: fmt.Sprintf(
						"[XDD-61] table %q column %q — matches sensitive pattern %q but has no -- @sensitive annotation",
						blk.tableName, colName, match),
					Advice: "Append annotations at the END of the column line, in a single comment. " +
						"Example: email VARCHAR(255) NOT NULL -- @sensitive. " +
						"Multiple: token_hash VARCHAR(255) NOT NULL -- @sensitive @archived. " +
						"Do NOT write on a separate line. Do NOT write -- @a -- @b.",
				})
			}
		}
	}
	return diags
}
