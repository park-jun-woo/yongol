//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what collectSensitiveKeys — DDL `-- @sensitive` 컬럼명을 sorted 리스트로 수집

package boot

import (
	"sort"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// collectSensitiveKeys walks all parsed DDL tables and returns a sorted, de-
// duplicated list of column names marked with `-- @sensitive`. The output
// feeds into blockLoggerInit which injects each name into the runtime
// sensitiveKeys map so ReplaceAttr masks them in slog output.
func collectSensitiveKeys(fs *yongol.Fullstack) []string {
	seen := make(map[string]bool)
	for _, t := range fs.DDLTables {
		collectSensitiveKeysFromTable(t.Columns, seen)
	}
	out := make([]string, 0, len(seen))
	for k := range seen {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
