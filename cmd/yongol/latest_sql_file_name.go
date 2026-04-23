//ff:func feature=cli type=util control=iteration dimension=1
//ff:what latestSQLFileName — 디렉토리 엔트리에서 문자열 순으로 가장 큰 *.sql 이름 선택
package main

import (
	"os"
	"strings"
)

// latestSQLFileName returns the lexicographically largest ".sql" filename
// among the given directory entries, or "" if none match. Migration files are
// prefixed with a zero-padded sequence so lexical order = chronological order.
func latestSQLFileName(entries []os.DirEntry) string {
	latest := ""
	for _, e := range entries {
		name := e.Name()
		if !strings.HasSuffix(name, ".sql") {
			continue
		}
		if name > latest {
			latest = name
		}
	}
	return latest
}
