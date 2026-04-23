//ff:func feature=migration type=parser control=iteration dimension=1
//ff:what BuildASTFromDir — <dir>/*.sql 파일을 읽어 Schema AST 조립 (skipFiles 제외)
package migration

import (
	"fmt"
	"os"
)

// BuildASTFromDir walks <dir> for *.sql files and returns a Schema.
// It intentionally re-reads the raw files rather than going through
// pkg/parser/ddl because that parser already lossily converts column
// types to Go types. For diff purposes we need the full PostgreSQL type
// information (INT vs BIGINT vs VARCHAR(255) …).
//
// skipFiles lists base file names to skip (e.g. ".generated_schema.sql").
func BuildASTFromDir(dir string, skipFiles []string) (*Schema, error) {
	s := NewSchema()
	files, err := listSQLFiles(dir, skipFiles)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", f, err)
		}
		if err := BuildASTFromSQL(s, string(data)); err != nil {
			return nil, fmt.Errorf("parse %s: %w", f, err)
		}
	}
	return s, nil
}
