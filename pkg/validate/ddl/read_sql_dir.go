//ff:func feature=validate type=util control=iteration dimension=1 topic=ddl-structural
//ff:what readSQLDir — read all .sql files in a directory and return them as a []sqlFile slice
package ddl

import (
	"os"
	"path/filepath"
	"strings"
)

func readSQLDir(dir string) []sqlFile {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}
	var out []sqlFile
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".sql") {
			continue
		}
		// Defensive: skip yongol's migration baseline even if it leaks
		// into the DDL directory. Phase010 (BUG-034) moved it to
		// arts/db/migrations/.latest_schema.sql; only pre-Phase010
		// leftover files would land here and generate removes them on
		// entry.
		if e.Name() == ".latest_schema.sql" || e.Name() == ".generated_schema.sql" {
			continue
		}
		full := filepath.Join(dir, e.Name())
		data, err := os.ReadFile(full)
		if err != nil {
			continue
		}
		out = append(out, sqlFile{path: full, name: e.Name(), content: string(data)})
	}
	return out
}
