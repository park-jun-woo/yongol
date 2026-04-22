//ff:func feature=validate type=util control=iteration dimension=1 topic=ddl-structural
//ff:what readSQLDir — 디렉토리 내 .sql 파일들을 sqlFile 슬라이스로 읽어 반환
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
		// Skip yongol's migration snapshot baseline — it is generated
		// output, not a DDL SSOT (see pkg/generate/migration).
		if e.Name() == ".generated_schema.sql" {
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
