//ff:func feature=validate type=rule control=iteration dimension=3 topic=ddl-structural
//ff:what D-8 — SERIAL/BIGSERIAL/SMALLSERIAL column types banned, replace with GENERATED ALWAYS AS IDENTITY

package ddl

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// d08SerialTypeBanned validates D-8: a column whose raw type token is
// SERIAL / BIGSERIAL / SMALLSERIAL is flagged as ERROR. The column must
// be rewritten with GENERATED ALWAYS AS IDENTITY.
func d08SerialTypeBanned(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	files := readDBSQLFiles(fs.SpecsDir)
	if len(files) == 0 {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, f := range files {
		for _, blk := range extractTableBlocks(f) {
			lines := strings.Split(blk.body, "\n")
			for offset, line := range lines {
				if d := d08CheckLine(f, blk, line, offset); d != nil {
					diags = append(diags, *d)
				}
			}
		}
	}
	return diags
}
