//ff:func feature=ddl type=util control=iteration dimension=1
//ff:what walkSQLFiles — dir 내 .sql 파일을 순회하며 각 파일 본문을 handler 로 전달
package ddl

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// walkSQLFiles reads every *.sql file directly under dir (non-recursive) and
// invokes handler(path, data) for each. Read errors are accumulated as
// PhaseParse/LevelError diagnostics, and diagnostics returned from handler
// are appended into the same slice. Returns a single combined diagnostics
// slice so call sites do not need to track dir/read/handler errors
// separately.
func walkSQLFiles(dir string, handler func(path string, data []byte) []diagnostic.Diagnostic) []diagnostic.Diagnostic {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return []diagnostic.Diagnostic{{
			File:    dir,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: "cannot read DDL directory: " + err.Error(),
		}}
	}
	var diags []diagnostic.Diagnostic
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".sql") {
			continue
		}
		// Defensive: skip yongol's migration baseline if it somehow
		// appears under the DDL directory. Phase010 (BUG-034) moved the
		// baseline to arts/db/migrations/.latest_schema.sql so this
		// guard now only covers leftover files from pre-Phase010 runs
		// (generate also removes them on entry).
		if e.Name() == ".latest_schema.sql" || e.Name() == ".generated_schema.sql" {
			continue
		}
		path := filepath.Join(dir, e.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			diags = append(diags, diagnostic.Diagnostic{
				File:    path,
				Line:    0,
				Phase:   diagnostic.PhaseParse,
				Level:   diagnostic.LevelError,
				Message: "cannot read SQL file: " + err.Error(),
			})
			continue
		}
		if more := handler(path, data); len(more) > 0 {
			diags = append(diags, more...)
		}
	}
	return diags
}
