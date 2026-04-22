//ff:func feature=validate type=util control=iteration dimension=1 topic=query-structural
//ff:what q01ScanForMissingName — SQL 파일을 훑어 `-- name:` 누락된 pending SQL line 감지

package query

import (
	"bufio"
	"os"
	"strings"
)

// q01ScanForMissingName opens file and scans lines looking for SQL statements
// that precede any `-- name:` header. Returns (firstPendingLine, true) when a
// bareword SQL line exists before the first name annotation; otherwise
// (0, false).
func q01ScanForMissingName(file string) (int, bool) {
	f, err := os.Open(file)
	if err != nil {
		return 0, false
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineNo := 0
	sawName := false
	pendingSQL := false
	pendingLine := 0
	for scanner.Scan() {
		lineNo++
		trimmed := strings.TrimSpace(scanner.Text())
		if trimmed == "" {
			continue
		}
		if strings.HasPrefix(trimmed, "-- name:") {
			sawName = true
			pendingSQL = false
			continue
		}
		if strings.HasPrefix(trimmed, "--") {
			continue
		}
		if !sawName && !pendingSQL {
			pendingSQL = true
			pendingLine = lineNo
		}
	}
	if pendingSQL && !sawName {
		return pendingLine, true
	}
	return 0, false
}
