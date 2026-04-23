//ff:func feature=migration type=util control=sequence
//ff:what shouldParseSQL — 디렉토리/비-SQL/skip 목록을 제외할지 결정
package migration

import "strings"

// shouldParseSQL reports whether a dir entry should be parsed as DDL.
// Rules: must be a file, must end with ".sql" (case-insensitive), must
// not appear in `skip`.
func shouldParseSQL(isDir bool, name string, skip map[string]bool) bool {
	if isDir {
		return false
	}
	if !strings.HasSuffix(strings.ToLower(name), ".sql") {
		return false
	}
	if skip[name] {
		return false
	}
	return true
}
