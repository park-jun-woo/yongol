//ff:func feature=validate type=util control=sequence topic=ddl-structural
//ff:what readDBSQLFiles — <specsDir>/db/*.sql 파일 목록을 non-recursive 읽기
package ddl

import "path/filepath"

// readDBSQLFiles reads <specsDir>/db/*.sql files (non-recursive).
// Returns nil if the directory is missing.
func readDBSQLFiles(specsDir string) []sqlFile {
	return readSQLDir(filepath.Join(specsDir, "db"))
}
