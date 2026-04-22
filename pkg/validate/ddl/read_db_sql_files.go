//ff:func feature=validate type=util control=sequence topic=ddl-structural
//ff:what readDBSQLFiles — read <specsDir>/db/*.sql files non-recursively
package ddl

import "path/filepath"

// readDBSQLFiles reads <specsDir>/db/*.sql files (non-recursive).
// Returns nil if the directory is missing.
func readDBSQLFiles(specsDir string) []sqlFile {
	return readSQLDir(filepath.Join(specsDir, "db"))
}
