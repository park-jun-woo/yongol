//ff:func feature=validate type=util control=iteration dimension=1 topic=query-structural
//ff:what ddlHasUUIDColumn — <specsDir>/db/*.sql 중 UUID 컬럼 선언이 하나라도 있으면 true

package query

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// uuidColumnRe matches a column definition whose declared PG type is UUID.
// The regex anchors on a column-name token followed by `UUID` as a whole
// word at the start of a stripped line. DDL parser (pkg/parser/ddl) reduces
// UUID to Go `string`, so Q-12 scans the raw .sql files directly.
var uuidColumnRe = regexp.MustCompile(`(?im)^\s*[A-Za-z_][A-Za-z0-9_]*\s+UUID\b`)

// ddlHasUUIDColumn returns true when any *.sql file in <specsDir>/db/
// declares at least one UUID-typed column. Used by Q-12 to gate the
// pgtype.UUID override check — projects without UUID columns must skip
// the rule entirely.
func ddlHasUUIDColumn(specsDir string) bool {
	dbDir := filepath.Join(specsDir, "db")
	entries, err := os.ReadDir(dbDir)
	if err != nil {
		return false
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".sql") {
			continue
		}
		body, err := os.ReadFile(filepath.Join(dbDir, e.Name()))
		if err != nil {
			continue
		}
		if uuidColumnRe.Match(body) {
			return true
		}
	}
	return false
}
